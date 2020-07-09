// Package yuml2svg generates yuml files using the yuml2svg Typescript package.
//
// Requires yarn to install yuml2svg globally.
// For more infos: https://github.com/aduh95/yuml2svg#readme
// For vscode snippets: https://github.com/aduh95/yuml2svg/blob/master/test/snippets.json
package yuml2svg

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"aduu.dev/utils/dash"
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

	// The runner to use. If it is not set a built-in runner is used.
	// This option is here for use in tests to record commands.
	r dash.Runner
}

var (
	errNoYumlSuffix = fmt.Errorf("err: does not have a yuml suffix")
)

func svgPath(yumlPath string) string {
	return strings.TrimSuffix(yumlPath, yumlSuffix) + svgSuffix
}

// GenerateYumlInplace generates an svg from the given yuml.
// The name is the source name with the extension exchanged with svg.
func GenerateYumlInplace(settings Settings, yumlPath string) (err error) {
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
		"svg-path", svg,
		"dark", settings.Dark)

	return settings.r.RunE(context.Background(),
		dash.TemplateSplitExpand(`yuml2svg {{- if .Dark}} --dark{{end}}`,
			struct {
				Dark bool
			}{
				Dark: settings.Dark,
			}),
		dash.WithStdinFile(yumlPath),
		dash.WithStdoutFile(svg),
	)
}

// Install yuml2svg with yarn.
func Install(r dash.Runner) (err error) {
	return r.RunE(context.Background(),
		dash.TemplateSplitExpand(`yarn global add yuml2svg`, ""))
}

// GenerateYumls walks the directory path and generates svgs from yuml files.
//
// The svg files are named the same as the yuml files with the suffix replaced.
func GenerateYumls(settings Settings, root string) (err error) {
	klog.InfoS("Generate yumls",
		"root", root,
		"dark", settings.Dark)

	// If the user did not set a runner then use the built-in.
	if settings.r == nil {
		settings.r = dash.NewRunner()
	}

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, yumlSuffix) || info.IsDir() {
			return nil
		}

		return GenerateYumlInplace(settings, path)
	})

	if err != nil {
		return err
	}

	return nil

}
