
```go
func (f Function) String() string {
	return fmt.Sprintf("\n```go\n%s```\n", f.Body)
}```

```go
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
}```

```go
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
}```

```go
func AutoMdDoc(inputFilePath string, outputFilePath string) {
	// goContent, err := ioutil.ReadFile(inputFilePath)
	content, err := io.ReadPath(inputFilePath)
	if err != nil {
		panic(err)
	}

	wrapper := detectContent(content)

	// err = ioutil.WriteFile(outputFilePath, wrapper, 0644)
	err = io.WritePath(outputFilePath, wrapper)
	if err != nil {
		panic(err)
	}
}```
