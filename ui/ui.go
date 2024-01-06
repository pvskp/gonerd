package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	title, desc string
}

func (i item) FilterValue() string {
	return ""
}

type Model struct {
	MarkedFiles  list.Model
	Downloadable list.Model
}

func NewModel() Model {
	items := []list.Item{
		item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{title: "Nutella", desc: "It's good on toast"},
		item{title: "Linux", desc: "Pretty much the best OS"},
		item{title: "Shampoo", desc: "Nothing like clean hair"},
		item{title: "Afternoon tea", desc: "Especially the tea sandwich part"},
		item{title: "Warm light", desc: "Like around 2700 Kelvin"},
		item{title: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
		item{title: "Gaffer’s tape", desc: "Basically sticky fabric"},
		item{title: "Terrycloth", desc: "In other words, towel fabric"},
	}

	return Model{
		MarkedFiles: list.New(items,
			list.NewDefaultDelegate(),
			20,
			20,
		),

		Downloadable: list.New(items,
			list.NewDefaultDelegate(),
			20,
			20,
		),
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.MarkedFiles.SetWidth(msg.Width)
		m.Downloadable.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.MarkedFiles, _ = m.MarkedFiles.Update(msg)
	m.Downloadable, cmd = m.Downloadable.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.Downloadable.View(), m.MarkedFiles.View())
}

func (m Model) Init() tea.Cmd { return nil }
