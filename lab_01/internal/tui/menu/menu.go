package menu

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Cursor int      // Местоположение курсора
	Items  []string // Пункты меню
}

type ItemSelectedMsg struct {
	Index int
}

func New(items []string) Model {
	return Model{
		Items: items,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.Cursor > 0 {
				m.Cursor--
			}
		case tea.KeyDown:
			if m.Cursor < len(m.Items)-1 {
				m.Cursor++
			}
		case tea.KeyEnter:
			return m, func() tea.Msg {
				return ItemSelectedMsg{Index: m.Cursor}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	var s string
	for i, item := range m.Items {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, item)
	}
	return lipgloss.NewStyle().
		Padding(1, 2).
		Render("Главное меню:\n\n" + s)
}

func (m Model) Init() tea.Cmd {
	return nil
}
