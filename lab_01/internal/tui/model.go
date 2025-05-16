package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	menuState state = iota
	regexInputState
	exprInputState
)

type LogicInterface interface {
	RegularToFAs(input string) (string, error)
	Emulate(input string) (string, error)
	//HandleExpression(regex, expr string) string
}
type AppModel struct {
	state      state
	menuCursor int
	regexInput textinput.Model
	exprInput  textinput.Model
	regex      string
	lastResult string
	logic      LogicInterface
}

func (m AppModel) Init() tea.Cmd {
	return nil
}
