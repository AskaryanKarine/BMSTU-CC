package commands

import (
	"bytes"
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/dfa"
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/fs"
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/nfa"
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/transform"
	"os"
	"os/exec"
)

const (
	nfaFileDot = "graphs/nfa.dot"
	nfaFilePng = "graphs/nfa.png"
	dfaFileDot = "graphs/dfa.dot"
	dfaFilePng = "graphs/dfa.png"
	minFileDot = "graphs/min.dot"
	minFilePng = "graphs/min.png"
	stepsDir   = "graphs/emulate/"
)

type Commands struct {
	outputBuffer bytes.Buffer
	dfa          *dfa.DFA
}

func (c *Commands) RegularToFAs(input string) (string, error) {
	c.outputBuffer.Reset()
	postfixForm := transform.InfixToPostfix(input)

	// Обработка NFA
	initialNFA := nfa.Build(postfixForm)
	if err := c.saveAutomaton("NFA", initialNFA.ToGraphviz(), nfaFileDot, nfaFilePng); err != nil {
		return c.output(), err
	}

	// Обработка DFA
	builtDFA := dfa.Build(initialNFA)
	if err := c.saveAutomaton("DFA", builtDFA.ToGraphviz(), dfaFileDot, dfaFilePng); err != nil {
		return c.output(), err
	}

	// Обработка Minimized DFA
	minimizeDFA := builtDFA.Minimize()
	if err := c.saveAutomaton("Minimized DFA", minimizeDFA.ToGraphviz(), minFileDot, minFilePng); err != nil {
		return c.output(), err
	}

	c.dfa = minimizeDFA

	return c.output(), nil
}

func (c *Commands) saveAutomaton(name, dotData, dotPath, pngPath string) error {
	//c.outputBuffer.WriteString(fmt.Sprintf("Сохраняется %s...\n", name))
	if err := os.WriteFile(dotPath, []byte(dotData), 0644); err != nil {
		msg := fmt.Sprintf("❌ ошибка при записи файла %s: %v\n", name, err)
		c.outputBuffer.WriteString(msg)

		return fmt.Errorf(msg)
	}

	cmd := exec.Command("dot", "-Tpng", "-o", pngPath, dotPath)
	if output, err := cmd.CombinedOutput(); err != nil {
		msg := fmt.Sprintf("❌ Ошибка Graphviz для %s: %v\nВывод: %s\n", name, err, string(output))
		c.outputBuffer.WriteString(msg)

		return fmt.Errorf(msg)
	}

	return nil
}

func (c *Commands) Emulate(input string) (string, error) {
	c.outputBuffer.Reset()

	err := fs.RecreateEmulateDir()
	if err != nil {
		msg := fmt.Sprintf("❌ ошибка при подготовке папки emulate: %v\n", err)
		c.outputBuffer.WriteString(msg)

		return c.output(), err
	}
	steps, ok := c.dfa.SimulateDFA(input)
	for i, step := range steps {
		name := fmt.Sprintf("step_%d", i)
		err := c.saveAutomaton(name, step, stepsDir+name+".dot", stepsDir+name+".png")
		if err != nil {
			return c.output(), err
		}
	}

	if ok {
		msg := fmt.Sprintf("✅ Строка %s допускается ДКА", input)
		c.outputBuffer.WriteString(msg)
	} else {
		msg := fmt.Sprintf("❌ Строка %s НЕ допускается ДКА", input)
		c.outputBuffer.WriteString(msg)
	}

	return c.output(), nil
}

func (c *Commands) output() string {
	return c.outputBuffer.String()
}
