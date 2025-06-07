package compiler

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	ast "github.com/AskaryanKarine/BMSTU-CC/cource/internal/tree"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/visitor"
	"github.com/antlr4-go/antlr/v4"
)

func Compiler(input, output string, isDOT bool) error {
	fileStream, err := antlr.NewFileStream(input)
	if err != nil {
		return fmt.Errorf("compiler input error: %w", err)
	}

	lexer := parser.NewKumirLexer(fileStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewKumirParser(tokens)

	tree := p.Program()

	if isDOT {
		err := ast.SaveTreeToFile(tree, p, output)
		if err != nil {
			return fmt.Errorf("ast build error: %w", err)
		}
	}

	v := visitor.NewIRVisitor()
	v.Visit(tree)

	//llFilename := output + ".ll"

	//err = v.SaveToFile(llFilename)
	//if err != nil {
	//	return fmt.Errorf("save llvm output: %w", err)
	//}

	//cmd := exec.Command("clang", llFilename, "-o", output, "-lm")
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//err = cmd.Run()
	//if err != nil {
	//	return fmt.Errorf("compile llvm to program: %w", err)
	//}

	fmt.Println(input, output)
	return nil
}
