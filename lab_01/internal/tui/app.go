package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewApp(logic LogicInterface) AppModel {
	regex := textinput.New()
	regex.Placeholder = "Например, (ab)*c"
	regex.PlaceholderStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true)
	regex.Prompt = "> "
	regex.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("201"))
	regex.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("230"))
	regex.Width = 50
	regex.Focus()

	expr := textinput.New()
	expr.Placeholder = "Например, ababc"
	expr.PlaceholderStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true)
	expr.Prompt = "> "
	expr.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("201"))
	expr.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("230"))
	expr.Width = 50

	return AppModel{
		state:      menuState,
		logic:      logic,
		regexInput: regex,
		exprInput:  expr,
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

		case regexInputState:
			return handleRegexInput(msg, m)

		case exprInputState:
			return handleExprInput(msg, m)
		}
	}
	return m, nil
}
