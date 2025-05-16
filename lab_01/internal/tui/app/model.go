package app

import (
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/tui/menu"
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/tui/textinput"
	_ "github.com/charmbracelet/bubbletea"
)

type state int

const (
	stateMenu state = iota
	stateRegexInput
	stateExprInput
)

// LogicInterface TODO: не забудь потом сменить!!!
type LogicInterface interface {
	HandleRegex(input string) string
	HandleExpression(regex, expr string) string
}

type Model struct {
	state       state
	menu        menu.Model
	regexInput  textinput.Model
	exprInput   textinput.Model
	regexResult string
	lastResult  string
	logic       LogicInterface
}
