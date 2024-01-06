package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pvskp/gonerd/ui"
)

func main() {
	model := ui.NewModel()
	model.Downloadable.Title = "Downloadable"
	model.Marked.Title = "MarkedFiles"
	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
