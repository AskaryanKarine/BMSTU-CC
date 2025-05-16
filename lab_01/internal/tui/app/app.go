package app

import (
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/tui/menu"
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/tui/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func Init(logic LogicInterface) Model {
	m := Model{
		menu: menu.New([]string{
			"Ввести регулярку",
			"Промоделировать выражение",
		}),
		regexInput: textinput.New("Введите регулярное выражение"),
		exprInput:  textinput.New("Введите выражение для проверки"),
		logic:      logic,
	}

	m.regexInput.OnSubmit = func(value string) tea.Cmd {
		m.regexInput.Input.Focus()
		m.regexResult = m.logic.HandleRegex(value)
		m.lastResult = m.regexResult
		return tea.Sequence(
			func() tea.Msg { return textinput.ToMenuMsg{} },
		)
	}

	m.exprInput.OnSubmit = func(value string) tea.Cmd {
		m.lastResult = m.logic.HandleExpression(m.regexResult, value)
		m.exprInput.Input.Focus()
		return tea.Sequence(
			func() tea.Msg { return textinput.ToMenuMsg{} },
		)
	}

	return m
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

	case menu.ItemSelectedMsg:
		return m.handleMenuSelection(msg)

	case textinput.ToMenuMsg:
		m.state = stateMenu
	}

	var cmd tea.Cmd

	switch m.state {
	case stateMenu:
		m.menu, cmd = m.menu.Update(msg)
		return m, cmd
	case stateRegexInput:
		m.regexInput, cmd = m.regexInput.Update(msg)
		return m, cmd
	case stateExprInput:
		m.exprInput, cmd = m.exprInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) handleMenuSelection(msg menu.ItemSelectedMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.Index {
	case 0: // Ввести регулярку
		m.state = stateRegexInput
		m.regexInput.Active = true
		cmd = m.regexInput.Input.Focus()

	case 1: // Промоделировать выражение
		if m.regexResult == "" {
			m.lastResult = "Сначала введите регулярное выражение!"
			return m, nil
		}
		m.state = stateExprInput
		m.exprInput.Active = true
		cmd = m.exprInput.Input.Focus()
	}

	return m, cmd
}

func (m Model) View() string {
	switch m.state {
	case stateMenu:
		return m.menu.View()
	case stateRegexInput:
		return m.regexInput.View()
	case stateExprInput:
		return m.exprInput.View()
	}
	return ""
}

func (m Model) Init() tea.Cmd {
	return nil
}
