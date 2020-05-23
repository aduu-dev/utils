package generate

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/ast/astutil"

	"aduu.dev/utils/exe"
)

const completionFile = "$HOME/go/pkg/mod/k8s.io/kubernetes@v1.14.0-beta.2/pkg/kubectl/cmd/completion/completion.go"
const outFilename = "kubernetes_completion.go"

//go:generate aduu generate

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func removeUnwantedDeclarations(settings settings) {
	src, err := ioutil.ReadFile(settings.To)
	if err != nil {
		fmt.Println(errors.WithMessage(err, "failed to read kubectl file "+settings.To))
		os.Exit(1)
	}

	// Create the AST by parsing src.

	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, outFilename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Delete unwanted.

	_ = astutil.Apply(f, func(c *astutil.Cursor) bool {
		n := c.Node()
		if d, ok := n.(*ast.FuncDecl); ok && (d.Name.Name == "NewCmdCompletion" || d.Name.Name == "RunCompletion") {
			c.Delete()
		}
		return true
	}, nil)

	// Create an ast.CommentMap from the ast.File's comments.
	// This helps keeping the association between comments
	// and AST nodes.

	cmap := ast.NewCommentMap(fset, f, f.Comments)

	// Remove the first variable declaration from the list of declarations.

	for i, decl := range f.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.VAR {
			copy(f.Decls[i:], f.Decls[i+1:])
			f.Decls = f.Decls[:len(f.Decls)-1]
		}
	}

	// Use the comment map to filter comments that don't belong anymore
	// (the comments associated with the variable declaration), and create
	// the new comments list.

	f.Comments = cmap.Filter(f).Comments()

	// Print the modified AST.

	var buf bytes.Buffer

	if err := format.Node(&buf, fset, f); err != nil {

		panic(err)

	}

	if err := ioutil.WriteFile(settings.To, buf.Bytes(), 0777); err != nil {
		handleError(err)
	}
}

type settings struct {
	From string
	To   string
}

var CMD = &cobra.Command{
	Use:   "generate",
	Short: "extracts completion file from kubectl.",
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(errors.WithMessage(errors.WithStack(err), "failed to get wd"))
			os.Exit(1)
		}

		settings := settings{
			completionFile,
			filepath.Join(wd, outFilename),
		}
		exe.Execute(`cp {{.From}} {{.To}}`, settings)

		handleError(os.Chmod(settings.To, 0755))

		removeUnwantedDeclarations(settings)

		exe.Execute(`goimports -w {{.To}}`, settings)
	},
}
