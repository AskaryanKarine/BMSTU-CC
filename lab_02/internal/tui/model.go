package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	menuState state = iota
	filePathInputState
)

const maxMenuStates = 3

type LogicInterface interface {
	LoadGrammar(input string) (string, error)
	EliminatingLeftRecursion() (string, error)
	EliminatingLeftIndirectRecursion() (string, error)
	EliminationUselessSymbols() (string, error)
}
type AppModel struct {
	state           state
	menuCursor      int
	filePathInput   textinput.Model
	initGrammarPath string
	outputResult    string
	logic           LogicInterface
}

func (m AppModel) Init() tea.Cmd {
	return nil
}
