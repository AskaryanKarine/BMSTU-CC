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
		if m.menuCursor < maxMenuStates {
			m.menuCursor++
		}
	case tea.KeyEnter:
		if m.menuCursor != 0 && m.initGrammarPath == "" {
			m.outputResult = "Сначала введите грамматику!"
			return m, nil
		}
		switch m.menuCursor {
		case 0:
			m.state = filePathInputState
			m.filePathInput.Focus()
			return m, textinput.Blink
		case 1:
			m.outputResult = processLogic(m.logic.EliminatingLeftRecursion)
			return m, nil
		case 2:
			m.outputResult = processLogic(m.logic.EliminatingLeftIndirectRecursion)
			return m, nil
		case 3:
			m.outputResult = processLogic(m.logic.EliminationUselessSymbols)
			return m, nil
		}
	}
	return m, nil
}

func handleGrammarInput(msg tea.KeyMsg, m AppModel) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		m.initGrammarPath = m.filePathInput.Value()
		result, err := m.logic.LoadGrammar(m.initGrammarPath)
		if err != nil {
			m.outputResult = errorStyle.Render(result)
		} else {
			m.outputResult = resultStyle.Render(result)
		}
		m.state = menuState
		m.filePathInput.Reset()
		return m, nil
	case tea.KeyEsc:
		m.state = menuState
		return m, nil
	}

	var cmd tea.Cmd
	m.filePathInput, cmd = m.filePathInput.Update(msg)
	return m, cmd
}

func processLogic(f func() (string, error)) string {
	result, err := f()
	if err != nil {
		return errorStyle.Render(result)
	} else {
		return resultStyle.Render(result)
	}
}
