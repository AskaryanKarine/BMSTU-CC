package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

// Стили для компонентов
var (
	appStyle = lipgloss.NewStyle().
			Margin(1, 2).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Foreground(lipgloss.Color("230"))

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("212")).
			MarginBottom(1).
			Align(lipgloss.Center)

	menuItemStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(lipgloss.Color("240"))

	activeMenuItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("212")).
				Bold(true)

	inputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1).
			Width(50)

	activeInputStyle = inputStyle.BorderForeground(lipgloss.Color("201"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Italic(true)

	resultStyle = lipgloss.NewStyle().
			PaddingTop(1).
			Foreground(lipgloss.Color("156")).
			MaxWidth(60)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true)
)

func (m AppModel) View() string {
	content := ""

	switch m.state {
	case menuState:
		content = renderMenu(m)
	case filePathInputState:
		content = renderGrammarInput(m)
	}

	return appStyle.Render(content)
}

func renderMenu(m AppModel) string {
	var menu string
	for i, item := range []string{"Ввести путь до файла с грамматикой", "Устранить левую рекурсию",
		"Устранить левую косвенную рекурсию", "Устранить бесполезные символы"} {
		style := menuItemStyle
		if m.menuCursor == i {
			style = activeMenuItemStyle
		}
		menu += style.Render(fmt.Sprintf("> %s", item)) + "\n"
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("Устранение различных штук"),
		menu,
		resultStyle.Render(m.outputResult),
		helpStyle.Render("\n↑/↓ - навигация • Enter - выбор • Ctrl+C - выход"),
	)
}

func renderGrammarInput(m AppModel) string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		titleStyle.Render("Введите путь до файла с грамматикой"),
		activeInputStyle.Render(m.filePathInput.View()),
		helpStyle.Render("\nEnter - подтвердить • Esc - отмена"),
	)
}
