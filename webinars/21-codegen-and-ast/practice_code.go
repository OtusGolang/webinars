package main

import (
	"fmt"
	"github.com/ettle/strcase"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

var test = `package main

import "fmt"

type MyStruct struct {
	MyName string 
	MyLogin string
	MyEmail string
}

func main() {
	fmt.Println("Hello")
}
`

func main() {
	f := token.NewFileSet()
	astFile, err := parser.ParseFile(f, "", test, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(astFile, func(n ast.Node) bool {
		d, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}

		if d.Tok != token.TYPE {
			return true
		}

		for _, ds := range d.Specs {
			t, ok := ds.(*ast.TypeSpec)
			if !ok {
				return true
			}

			structType, ok := t.Type.(*ast.StructType)
			if !ok {
				return true
			}
			structName := t.Name.Name

			for _, field := range structType.Fields.List {
				if field.Tag != nil {
					continue
				}

				field.Tag = &ast.BasicLit{}
				field.Tag.Value = fmt.Sprintf("`json:\"%s,omitempty\"`", strcase.ToSnake(field.Names[0].Name))
			}

			if d.Doc.Text() != "" {
				return true
			}
			d.Doc = &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text:  fmt.Sprintf("// %s ... please documentation", structName),
						Slash: structType.Struct,
					},
				},
			}
		}

		return true
	})

	//ast.Print(f, astFile)
	err = printer.Fprint(os.Stdout, f, astFile)
	if err != nil {
		panic(err)
	}
}
