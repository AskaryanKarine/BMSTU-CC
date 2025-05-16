package main

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/commands"
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/fs"
	"github.com/AskaryanKarine/BMSTU-CC/lab_01/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	err := fs.DeleteGraphsDir()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	err = fs.CreateGraphsDir()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	logic := &commands.Commands{}
	p := tea.NewProgram(tui.NewApp(logic))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	//logic.RegularToFAs("(ab)*с")
}
