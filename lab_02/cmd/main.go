package main

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/lab_02/internal/commands"
	"github.com/AskaryanKarine/BMSTU-CC/lab_02/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	logic := &commands.Command{}
	p := tea.NewProgram(tui.NewApp(logic))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
