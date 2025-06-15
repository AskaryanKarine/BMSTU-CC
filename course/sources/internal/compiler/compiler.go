package compiler

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	ast "github.com/AskaryanKarine/BMSTU-CC/cource/internal/tree"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/visitor"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"os/exec"
)

func Compiler(input, output string, isDOT bool) error {
	fileStream, err := antlr.NewFileStream(input)
	if err != nil {
		return fmt.Errorf("compiler input error: %w", err)
	}

	lexer := parser.NewKumirLexer(fileStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewKumirParser(tokens)

	errListener := newKumirErrorListener()
	p.AddErrorListener(errListener)
	if len(errListener.errs) > 0 {
		return fmt.Errorf("compiler parser error: %w", errListener.errs[0])
	}

	tree := p.Program()

	if isDOT {
		err := ast.SaveTreeToFile(tree, p, output)
		if err != nil {
			return fmt.Errorf("ast build error: %w", err)
		}
	}

	v := visitor.NewIRVisitor()
	v.Visit(tree)

	if len(v.Errors) > 0 {
		return fmt.Errorf("compiler errors: %v", v.Errors)
	}

	llFilename := output + ".ll"

	err = v.WriteToFile(llFilename)
	if err != nil {
		return fmt.Errorf("save llvm output: %w", err)
	}

	cmd := exec.Command("clang", llFilename, "-o", output, "-lm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("compile llvm to program: %w", err)
	}

	return nil
}

type customErrorListener struct {
	*antlr.DefaultErrorListener
	IsError bool
}

type kumirErrorListener struct {
	*antlr.DefaultErrorListener
	errs []error
}

func newKumirErrorListener() *kumirErrorListener {
	return &kumirErrorListener{
		DefaultErrorListener: new(antlr.DefaultErrorListener),
		errs:                 make([]error, 0),
	}
}

func (l *kumirErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.errs = append(l.errs, fmt.Errorf("syntax error at %d:%d – %s", line, column, msg))
}
