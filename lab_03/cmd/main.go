package main

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/lab_03/internal/fs"
	"github.com/AskaryanKarine/BMSTU-CC/lab_03/internal/lexer"
	"github.com/AskaryanKarine/BMSTU-CC/lab_03/internal/parser"
	"os"
)

const (
	dotFile = "ast.dot"
	pngFile = "ast.png"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Используйте: go run [cmd/]main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	input, err := fs.ReadProgramFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lex := lexer.NewLexer(input)
	tokens := lex.Tokenize()

	pars := parser.NewParser(tokens)
	ast, ok := pars.Parse()

	if ok {
		fmt.Println("Успешно обработано")

		if err := fs.SaveDOTToFile(parser.GenerateDOT(ast), dotFile, pngFile); err != nil {
			fmt.Printf("Ошибка сохранения DOT файла: %v\n", err)
			return
		}
		fmt.Printf("AST сохранено: %s, %s\n", dotFile, pngFile)
	} else {
		// Вывод информации об ошибке
		tok := pars.CurrentToken()
		fmt.Printf("Неожиданный токен: %s\n", tok.Literal)

	}
}
