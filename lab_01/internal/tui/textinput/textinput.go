package textinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ToMenuMsg struct{}

type Model struct {
	Input    textinput.Model // Готовый компонент
	Active   bool
	OnSubmit func(string) tea.Cmd
	OnCancel tea.Cmd
}

func New(placeholder string) Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()

	return Model{
		Input:  ti,
		Active: false,
		OnCancel: tea.Sequence(
			func() tea.Msg { return ToMenuMsg{} }, textinput.Blink,
		),
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.Active {
		return m, nil
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.OnSubmit != nil {
				return m, m.OnSubmit(m.Input.Value())
			}
		case tea.KeyEsc:
			if m.OnCancel != nil {
				return m, m.OnCancel
			}
		}
	}

	m.Input, cmd = m.Input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if !m.Active {
		return ""
	}
	return m.Input.View() + "\n(esc чтобы вернуться)"
}
