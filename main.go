package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pvskp/gonerd/ui"
)

func main() {
	if _, err := tea.NewProgram(ui.NewModel(), tea.WithAltScreen()).Run(); err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
