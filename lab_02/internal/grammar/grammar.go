package grammar

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Empty = "eps"
)

type Grammar struct {
	NonTerminals []string
	Terminals    []string
	Start        string
	Rules        map[string][][]string
}

func ReadGrammarFromFile(filename string) (*Grammar, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	g := &Grammar{Rules: make(map[string][][]string)}

	if !readLine(scanner) {
		return nil, fmt.Errorf("неожиданный конец файла")
	}
	if g.NonTerminals, err = readSymbols(scanner); err != nil {
		return nil, fmt.Errorf("нетерминал: %w", err)
	}

	if !readLine(scanner) {
		return nil, fmt.Errorf("неожиданный конец файла")
	}
	if g.Terminals, err = readSymbols(scanner); err != nil {
		return nil, fmt.Errorf("терминал: %w", err)
	}

	// Read rules
	if !readLine(scanner) {
		return nil, fmt.Errorf("неожиданный конец файла")
	}
	nRules, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, fmt.Errorf("некорректное количество правил: %w", err)
	}

	for i := 0; i < nRules; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf("правило %d: неожиданный конец файла", i+1)
		}
		line := strings.ReplaceAll(scanner.Text(), " ", "")
		parts := strings.Split(line, "->")
		if len(parts) != 2 {
			return nil, fmt.Errorf("некорректный формат правил: %s", line)
		}

		lhs := parts[0]
		rhs := []string{Empty}
		if parts[1] != Empty {
			rhs = splitSymbols(parts[1])
		}
		g.Rules[lhs] = append(g.Rules[lhs], rhs)
	}

	if !readLine(scanner) {
		return nil, fmt.Errorf("неожиданный конец файла")
	}
	g.Start = strings.TrimSpace(scanner.Text())

	return g, nil
}

func readLine(scanner *bufio.Scanner) bool {
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) != "" {
			return true
		}
	}
	return false
}

func readSymbols(scanner *bufio.Scanner) ([]string, error) {
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, fmt.Errorf("некорретный счетчик: %w", err)
	}

	if !scanner.Scan() {
		return nil, fmt.Errorf("неожиданный конец файла")
	}

	symbols := strings.Fields(scanner.Text())
	if len(symbols) != n {
		return nil, fmt.Errorf("несовпадение количества: ожидалось %d, прочитано %d", n, len(symbols))
	}

	return symbols, nil
}

func splitSymbols(s string) []string {
	var symbols []string
	for i := 0; i < len(s); {
		if i+1 < len(s) && s[i+1] == '\'' {
			symbols = append(symbols, s[i:i+2])
			i += 2
		} else {
			symbols = append(symbols, string(s[i]))
			i++
		}
	}
	return symbols
}

func (g *Grammar) String() string {
	result := ""
	result += fmt.Sprintf("%d\n", len(g.NonTerminals))
	result += strings.Join(g.NonTerminals, " ") + "\n"

	result += fmt.Sprintf("%d\n", len(g.Terminals))
	result += strings.Join(g.Terminals, " ") + "\n"

	cntRules := 0
	for _, rules := range g.Rules {
		cntRules += len(rules)
	}

	result += fmt.Sprintf("%d\n", cntRules)
	startsRules := g.Rules[g.Start]
	for _, rules := range startsRules {
		result += fmt.Sprintf("%s -> %s\n", g.Start, strings.Join(rules, " "))
	}

	for nt, rules := range g.Rules {
		if nt == g.Start {
			continue
		}
		for _, rule := range rules {
			result += fmt.Sprintf("%s -> %s\n", nt, strings.Join(rule, " "))
		}
	}

	result += fmt.Sprintf("%s", g.Start)

	return result
}

func (g *Grammar) Copy() *Grammar {
	newG := &Grammar{
		NonTerminals: make([]string, len(g.NonTerminals)),
		Terminals:    make([]string, len(g.Terminals)),
		Start:        g.Start,
		Rules:        make(map[string][][]string),
	}
	copy(newG.NonTerminals, g.NonTerminals)
	copy(newG.Terminals, g.Terminals)
	for nt, productions := range g.Rules {
		// Создаем новый слайс для продукций
		copiedProductions := make([][]string, len(productions))

		for i, prod := range productions {
			// Копируем каждую продукцию
			copiedProductions[i] = make([]string, len(prod))
			copy(copiedProductions[i], prod)
		}

		newG.Rules[nt] = copiedProductions
	}

	return newG
}
