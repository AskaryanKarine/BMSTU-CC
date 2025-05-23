package commands

import (
	"bytes"
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/lab_02/internal/fs"
	"github.com/AskaryanKarine/BMSTU-CC/lab_02/internal/grammar"
)

const (
	leftRecSuf      = "_lr"
	leftIndirectSuf = "_ilr"
	UnlessSymSuf    = "_us"
)

type Command struct {
	grammarFileName string
	grammar         *grammar.Grammar
	outputBuffer    bytes.Buffer
}

func (c *Command) output() string {
	return c.outputBuffer.String()
}

func (c *Command) LoadGrammar(input string) (string, error) {
	c.outputBuffer.Reset()
	gram, err := grammar.ReadGrammarFromFile(input)
	if err != nil {
		msg := fmt.Sprintf("ошибка при чтении файла %s: %v\n", input, err)
		c.outputBuffer.WriteString(msg)
		return c.output(), err
	}
	msg := fmt.Sprintf("Грамматика %s прочитана успешно", input)
	c.outputBuffer.WriteString(msg)
	c.grammar = gram
	c.grammarFileName = input
	return c.output(), nil
}

func (c *Command) EliminatingLeftRecursion() (string, error) {
	c.outputBuffer.Reset()
	c.outputBuffer.WriteString("Вызов устранения левой рекурсии\n")
	out := c.grammar.Copy().EliminateLeftRecursion().String()

	filename := fs.AddSuffixToFilename(c.grammarFileName, leftRecSuf)
	err := fs.WriteStringToFile(out, filename)
	if err != nil {
		msg := fmt.Sprintf("Ошибка при записи файла %s: %v\n", filename, err)
		c.outputBuffer.WriteString(msg)

		return c.output(), err
	}

	msg := fmt.Sprintf("Грамматика записана в файл %s\n", filename)
	c.outputBuffer.WriteString(msg)

	return c.output(), nil
}

func (c *Command) EliminatingLeftIndirectRecursion() (string, error) {
	c.outputBuffer.Reset()
	c.outputBuffer.WriteString("Вызов устранения косвенной левой рекурсии\n")

	out := c.grammar.Copy().RemoveCycles().EliminateLeftRecursion().String()
	filename := fs.AddSuffixToFilename(c.grammarFileName, leftIndirectSuf)
	err := fs.WriteStringToFile(out, filename)
	if err != nil {
		msg := fmt.Sprintf("Ошибка при записи файла %s: %v\n", filename, err)
		c.outputBuffer.WriteString(msg)

		return c.output(), err
	}

	msg := fmt.Sprintf("Грамматика записана в файл %s\n", filename)
	c.outputBuffer.WriteString(msg)

	return c.output(), nil
}

func (c *Command) EliminationUselessSymbols() (string, error) {
	c.outputBuffer.Reset()
	c.outputBuffer.WriteString("Вызов устранения бесполезных символов\n")
	out := c.grammar.Copy().EliminationUselessSymbols().String()

	filename := fs.AddSuffixToFilename(c.grammarFileName, UnlessSymSuf)
	err := fs.WriteStringToFile(out, filename)
	if err != nil {
		msg := fmt.Sprintf("Ошибка при записи файла %s: %v\n", filename, err)
		c.outputBuffer.WriteString(msg)

		return c.output(), err
	}

	msg := fmt.Sprintf("Грамматика записана в файл %s\n", filename)
	c.outputBuffer.WriteString(msg)

	return c.output(), nil
}
