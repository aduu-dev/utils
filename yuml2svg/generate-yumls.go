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

// Settings contain options used during yuml->svg conversion.
type Settings struct {
	// If Dark is true then the svg is rendered for dark mode, edges are white.
	// If Dark is false then the svg is rendered for light mode, edges are black.
	Dark bool

	// The shell to use. If it is not set a built-in shell is used.
	// This option is here for running tests.
	shell shell.Shell
}

var (
	errNoYumlSuffix = fmt.Errorf("err: does not have a yuml suffix")
)

func svgPath(yumlPath string) string {
	return strings.TrimSuffix(yumlPath, yumlSuffix) + svgSuffix
}

func generateYumlInplace(settings Settings, yumlPath string) (err error) {
	svg := svgPath(yumlPath)

	return GenerateYuml(settings, yumlPath, svg)
}

// GenerateYuml generates an svg from a given yuml file.
func GenerateYuml(settings Settings, yumlPath string, toSvgPath string) (err error) {
	if !strings.HasSuffix(yumlPath, yumlSuffix) {
		klog.ErrorS(errNoYumlSuffix, "does not have a yuml suffix",
			"yuml-path", yumlPath)

		return fmt.Errorf("does not have yuml suffix")
	}

	svg := toSvgPath
	klog.InfoS("Generating yuml",
		"yuml-path", yumlPath,
		"svg-path", svg)

	darkModeOption := "--dark "
	if !settings.Dark {
		darkModeOption = ""
	}

	settings.shell.RunScript(fmt.Sprintf(`cat %s | yuml2svg %s> %s`, yumlPath, darkModeOption, svg))

	return settings.shell.Err()
}

// Install yuml2svg with yarn.
func Install(sh shell.Shell) {
	sh.RunScript(`yarn global add yuml2svg`)
}

// GenerateYumls walks the directory path and generates svgs from yuml files.
//
// The svg files are named the same as the yuml files with the suffix replaced.
func GenerateYumls(settings Settings, root string) (err error) {
	klog.InfoS("Generate yumls",
		"root", root)

	// If the user did not set a shell then use the built-in.
	if settings.shell == nil {
		settings.shell = shell.New()
	}

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, yumlSuffix) || info.IsDir() {
			return nil
		}

		return generateYumlInplace(settings, path)
	})

	if err != nil {
		return err
	}

	return nil

}
