// Package yuml2svg generates yuml files using the yuml2svg Typescript package. Requires yarn to install yuml2svg globally.
package yuml2svg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"aduu.dev/utils/shell"
	"k8s.io/klog/v2"
)

const (
	yumlSuffix = ".yuml"
	svgSuffix  = ".svg"
)

var (
	errNoYumlSuffix = fmt.Errorf("err: does not have a yuml suffix")
)

func svgPath(yumlPath string) string {
	return strings.TrimSuffix(yumlPath, yumlSuffix) + svgSuffix
}

func generateYuml(sh shell.Shell, yumlPath string) (err error) {
	if !strings.HasSuffix(yumlPath, yumlSuffix) {
		klog.ErrorS(errNoYumlSuffix, "does not have a yuml suffix",
			"yuml-path", yumlPath)

		return fmt.Errorf("does not have yuml suffix")
	}

	svg := svgPath(yumlPath)
	klog.InfoS("Generating yuml",
		"yuml-path", yumlPath,
		"svg-path", svg)

	sh.RunScript(fmt.Sprintf(`cat %s | yuml2svg --dark > %s`, yumlPath, svg))

	return sh.Err()
}

// Install yuml2svg with yarn.
func Install(sh shell.Shell) {
	sh.RunScript(`yarn global add yuml2svg`)
}

// GenerateYumls walks the directory path and generates svgs from yuml files.
//
// The svg files are named the same as the yuml files with the suffix replaced.
func GenerateYumls(root string) (err error) {
	klog.InfoS("Generate yumls",
		"root", root)

	sh := shell.New()

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, yumlSuffix) || info.IsDir() {
			return nil
		}

		return generateYuml(sh, path)
	})

	if err != nil {
		return err
	}

	return nil

}
