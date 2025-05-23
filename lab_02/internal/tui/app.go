package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewApp(logic LogicInterface) AppModel {
	grammar := textinput.New()
	grammar.Placeholder = "Например, ./data/g1.txt"
	grammar.PlaceholderStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true)
	grammar.Prompt = "> "
	grammar.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("201"))
	grammar.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("230"))
	grammar.Width = 50
	grammar.Focus()

	return AppModel{
		state:         menuState,
		logic:         logic,
		filePathInput: grammar,
	}
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

		switch m.state {
		case menuState:
			return handleMenuUpdate(msg, m)

		case filePathInputState:
			return handleGrammarInput(msg, m)
		}
	}
	return m, nil
}
