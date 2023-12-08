package helper

import (
	"fmt"

	"go/ast"
	"go/parser"
	"go/token"

	"github.com/thuongtruong1009/gouse/io"
	"github.com/thuongtruong1009/gouse/types"
)

type Function struct {
	Name string
	Body string
}

func (f Function) String() string {
	return fmt.Sprintf("\n```go\n%s```\n", f.Body)
}

func detectContent(content []byte) []byte {
	var result []string
	functions := extractFunctions(content)
	for _, function := range functions {
		tmpFunc := &Function{
			Name: function.Name,
			Body: function.Body,
		}
		result = append(result, fmt.Sprintf("%s", tmpFunc.String()))
	}

	return types.StringsToBytes(result)
}

func extractFunctions(code []byte) []Function {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "example.go", string(code), parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return nil
	}

	var functions []Function

	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			funcName := fn.Name.Name
			funcBody := string(code[fn.Pos()-1 : fn.End()])

			functions = append(functions, Function{
				Name: funcName,
				Body: funcBody,
			})
		}
	}

	return functions
}

func AutoMdDoc(inputFilePath string, outputFilePath string) {
	content, err := io.ReadPath(inputFilePath)
	if err != nil {
		panic(err)
	}

	wrapper := detectContent(content)

	err = io.WritePath(outputFilePath, wrapper)
	if err != nil {
		panic(err)
	}
}
