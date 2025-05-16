package dfa

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/nfa"
)

type State struct {
	ID          int
	NFAStates   map[int]bool
	Transitions map[rune]int
	IsFinal     bool
}

type DFA struct {
	Start    int
	States   map[int]*State
	Alphabet []rune
}

func NewState(id int, nfaStates map[int]bool, isFinal bool) *State {
	return &State{
		ID:          id,
		NFAStates:   nfaStates,
		Transitions: make(map[rune]int),
		IsFinal:     isFinal,
	}
}

func Build(nfa *nfa.NFA) *DFA {
	alphabet := nfa.ExtractAlphabet()

	dfa := &DFA{
		States:   make(map[int]*State),
		Alphabet: alphabet,
	}

	startedStates := make(map[int]bool)
	startedStates[nfa.Start.ID] = true

	startNFAStates := nfa.EpsilonClosure(startedStates)
	dfa.Start = 0
	dfa.States[0] = NewState(0, startNFAStates, nfa.IsFinalState(startNFAStates))

	queue := []int{0}
	processed := make(map[int]bool)

	stateID := 1

	for len(queue) > 0 {
		currentStateID := queue[0]
		queue = queue[1:]

		if processed[currentStateID] {
			continue
		}
		processed[currentStateID] = true

		currentState := dfa.States[currentStateID]

		for _, symbol := range alphabet {
			nextNFAStates := make(map[int]bool)
			for nfaStateID := range currentState.NFAStates {
				state := nfa.StateByID(nfaStateID)
				for _, nextState := range state.Transitions[symbol] {
					nextNFAStates[nextState.ID] = true
				}
			}

			nextNFAStates = nfa.EpsilonClosure(nextNFAStates)

			if len(nextNFAStates) == 0 {
				continue
			}

			found := false
			var nextStateID int
			for id, state := range dfa.States {
				if statesEqual(state.NFAStates, nextNFAStates) {
					found = true
					nextStateID = id
					break
				}
			}

			if !found {
				nextStateID = stateID
				dfa.States[nextStateID] = NewState(nextStateID, nextNFAStates, nfa.IsFinalState(nextNFAStates))
				queue = append(queue, nextStateID)
				stateID++
			}

			currentState.Transitions[symbol] = nextStateID
		}
	}

	return dfa
}

func statesEqual(a, b map[int]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for stateID := range a {
		if !b[stateID] {
			return false
		}
	}
	return true
}

func (dfa *DFA) ToGraphviz() string {
	graph := "digraph DFA {\n"
	graph += "  rankdir=LR;\n"
	graph += "  node [shape = circle];\n"

	graph += fmt.Sprintf("  start [shape = point];\n")
	graph += fmt.Sprintf("  start -> %d;\n", dfa.Start)

	for _, state := range dfa.States {
		if state.IsFinal {
			graph += fmt.Sprintf("  %d [shape = doublecircle];\n", state.ID)
		} else {
			graph += fmt.Sprintf("  %d [shape = circle];\n", state.ID)
		}
	}

	for _, state := range dfa.States {
		for symbol, nextStateID := range state.Transitions {
			graph += fmt.Sprintf("  %d -> %d [label=\"%c\"];\n", state.ID, nextStateID, symbol)
		}
	}

	graph += "}\n"
	return graph
}

func (dfa *DFA) SimulateDFA(input string) ([]string, bool) {
	var steps []string
	currentStateID := dfa.Start
	currentState := dfa.States[currentStateID]

	steps = append(steps, dfa.ToGraphvizWithHighlight(currentStateID, "Начало"))

	for i, symbol := range input {
		if nextStateID, exists := currentState.Transitions[symbol]; exists {
			currentStateID = nextStateID
			currentState = dfa.States[currentStateID]
			steps = append(steps, dfa.ToGraphvizWithHighlight(currentStateID, fmt.Sprintf("Шаг %d -- символ '%c'", i+1, symbol)))
		} else {
			steps = append(steps, dfa.ToGraphvizWithError(currentStateID, symbol))
			return steps, false
		}
	}

	isAccepted := currentState.IsFinal
	if isAccepted {
		steps = append(steps, dfa.ToGraphvizWithHighlight(currentStateID, "Допускается"))
	} else {
		steps = append(steps, dfa.ToGraphvizWithHighlight(currentStateID, "НЕ допускается"))
	}

	return steps, isAccepted
}

func (dfa *DFA) ToGraphvizWithHighlight(currentStateID int, description string) string {
	graph := "digraph DFA {\n"
	graph += "  rankdir=LR;\n"
	graph += "  node [shape = circle];\n"

	graph += fmt.Sprintf("  start [shape = point];\n")
	graph += fmt.Sprintf("  start -> %d;\n", dfa.Start)

	for _, state := range dfa.States {
		if state.IsFinal {
			graph += fmt.Sprintf("  %d [shape = doublecircle];\n", state.ID)
		} else {
			graph += fmt.Sprintf("  %d [shape = circle];\n", state.ID)
		}
	}

	graph += fmt.Sprintf("  %d [color=red, fontcolor=red];\n", currentStateID)

	graph += fmt.Sprintf("  labelloc=\"t\";\n")
	graph += fmt.Sprintf("  label=\"%s\";\n", description)

	for _, state := range dfa.States {
		for symbol, nextStateID := range state.Transitions {
			graph += fmt.Sprintf("  %d -> %d [label=\"%c\"];\n", state.ID, nextStateID, symbol)
		}
	}

	graph += "}\n"
	return graph
}

func (dfa *DFA) ToGraphvizWithError(currentStateID int, symbol rune) string {
	graph := "digraph DFA {\n"
	graph += "  rankdir=LR;\n"
	graph += "  node [shape = circle];\n"

	graph += fmt.Sprintf("  start [shape = point];\n")
	graph += fmt.Sprintf("  start -> %d;\n", dfa.Start)

	for _, state := range dfa.States {
		if state.IsFinal {
			graph += fmt.Sprintf("  %d [shape = doublecircle];\n", state.ID)
		} else {
			graph += fmt.Sprintf("  %d [shape = circle];\n", state.ID)
		}
	}

	graph += fmt.Sprintf("  %d [color=red, fontcolor=red];\n", currentStateID)
	graph += fmt.Sprintf("  %d -> ошибка [label=\"%c\"];\n", currentStateID, symbol)
	graph += "  ошибка [shape=box, color=red, fontcolor=red];\n"

	graph += fmt.Sprintf("  labelloc=\"t\";\n")
	graph += fmt.Sprintf("  label=\"Ошибка: нет перехода для символа '%c'\";\n", symbol)

	for _, state := range dfa.States {
		for symbol, nextStateID := range state.Transitions {
			graph += fmt.Sprintf("  %d -> %d [label=\"%c\"];\n", state.ID, nextStateID, symbol)
		}
	}

	graph += "}\n"
	return graph
}
