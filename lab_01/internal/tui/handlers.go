package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func handleMenuUpdate(msg tea.KeyMsg, m AppModel) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyUp:
		if m.menuCursor > 0 {
			m.menuCursor--
		}
	case tea.KeyDown:
		if m.menuCursor < 1 {
			m.menuCursor++
		}
	case tea.KeyEnter:
		switch m.menuCursor {
		case 0:
			m.state = regexInputState
			m.regexInput.Focus()
			return m, textinput.Blink
		case 1:
			if m.regex == "" {
				m.lastResult = "Сначала введите регулярное выражение!"
				return m, nil
			}
			m.state = exprInputState
			m.exprInput.Focus()
			return m, textinput.Blink
		}
	}
	return m, nil
}

func handleRegexInput(msg tea.KeyMsg, m AppModel) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		m.regex = m.regexInput.Value()
		result, err := m.logic.RegularToFAs(m.regex)
		if err != nil {
			m.lastResult = errorStyle.Render(err.Error())
		} else {
			m.lastResult = resultStyle.Render(result)
		}
		m.state = menuState
		m.regexInput.Reset()
		return m, nil
	case tea.KeyEsc:
		m.state = menuState
		return m, nil
	}

	var cmd tea.Cmd
	m.regexInput, cmd = m.regexInput.Update(msg)
	return m, cmd
}

func handleExprInput(msg tea.KeyMsg, m AppModel) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		result, err := m.logic.Emulate(m.exprInput.Value())
		if err != nil {
			m.lastResult = errorStyle.Render(err.Error())
		} else {
			m.lastResult = resultStyle.Render(result)
		}
		m.state = menuState
		m.exprInput.Reset()
		return m, nil
	case tea.KeyEsc:
		m.state = menuState
		return m, nil
	}

	var cmd tea.Cmd
	m.exprInput, cmd = m.exprInput.Update(msg)
	return m, cmd
}
