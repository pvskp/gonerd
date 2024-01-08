package ui

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pvskp/gonerd/adapters"
	"github.com/pvskp/gonerd/events"
	"github.com/pvskp/gonerd/structs"
)

var (
	titleStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("205")).
			Align(lipgloss.Center)
		// MarginLeft(15)

	docStyle = lipgloss.NewStyle().
			Margin(1, 1).
			Border(lipgloss.NormalBorder())

	markedListStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Align(lipgloss.Left)

	downloadableListStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), true, true, true, true).
				Align(lipgloss.Left)
)

const (
	DownloadablePercentageSize = 0.7
	MarkedPercentageSize       = 0.3

	stateDownloadable = iota
	stateMarked
	totalStates
)

type Model struct {
	Marked       list.Model
	Downloadable list.Model
	state        int
	width        int
	height       int
}

func NewModel() (m *Model) {
	d := []list.Item{
		structs.NewItem("Fira Code"),
		structs.NewItem("Anonymous Pro"),
		structs.NewItem("Inconsolata"),
	}

	mr := []list.Item{
		structs.NewItem("Hack"),
		structs.NewItem("Hack"),
		structs.NewItem("Hack"),
	}

	m = &Model{
		Marked: list.New(
			mr,
			structs.ItemDelegate{},
			0,
			0,
		),

		Downloadable: list.New(
			d,
			structs.ItemDelegate{},
			0,
			0,
		),
		state: stateDownloadable,
	}

	m.Downloadable.Title = "Downloadable"
	m.Downloadable.Styles.Title = titleStyle
	m.Downloadable.SetShowHelp(false)
	m.Downloadable.SetShowStatusBar(false)

	m.Marked.Title = "Marked"
	m.Marked.Styles.Title = titleStyle
	m.Marked.SetShowHelp(false)
	m.Marked.SetShowStatusBar(false)

	log.Printf("[]list.Items = %v\n", m.Downloadable.Items())

	for _, v := range m.Downloadable.Items() {
		if item, ok := v.(structs.Item); ok {
			s := item.Content
			log.Printf("curr item = %v\n", s)
			log.Printf("curr item type = %T\n", s)
		}
	}
	return
}

func (m *Model) changeState() {
	m.state = (m.state + 1) % totalStates
}

func (m *Model) moveLeft() {
	if m.state-1 >= 0 {
		m.state--
	}
}

func (m *Model) moveRight() {
	if m.state+1 < totalStates {
		m.state++
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case events.FontDownloadedMsg:
		log.Println("FontDownloadedMsg received")

		// if msg.Success {
		// 	m.Marked.RemoveItem(m.Marked.Index())
		// }

		// if msg.Index == len(m.Marked.Items()) {
		// }

		// if len(m.Marked.Items()) == 0 {
		// 	m.changeState()
		// }

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.width = msg.Width
		m.height = msg.Height

		setDownloadableSize(&m, m.width-h, msg.Height-v)
		setMarkedSize(&m, m.width-h, msg.Height-v)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.changeState()
		case "l":
			m.moveRight()
		case "h":
			m.moveLeft()
		}

		switch m.state {
		case stateDownloadable:
			switch msg.String() {
			case "enter":
				m.Marked.InsertItem(
					len(m.Marked.Items()),
					m.Downloadable.SelectedItem(),
				)
			}

		case stateMarked:
			switch msg.String() {
			case "d":
				m.Marked.RemoveItem(m.Marked.Index())
			case "i":
				items := m.Marked.Items()
				if len(m.Marked.Items()) > 0 {
					adapters.DownloadSingleFont(items[m.Marked.Index()], m.Marked.Index())
				}
			}
		}
	}

	switch m.state {
	case stateDownloadable:
		m.Downloadable, cmd = m.Downloadable.Update(msg)
		cmds = append(cmds, cmd)
	case stateMarked:
		m.Marked, cmd = m.Marked.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func setDownloadableSize(m *Model, width, height int) {
	w := int(float64(width) * DownloadablePercentageSize)
	downloadableListStyle.Width(w)
	downloadableListStyle.Height(height)
	m.Downloadable.SetSize(w, height)
}

func setMarkedSize(m *Model, width, height int) {
	w := int(float64(width) * MarkedPercentageSize)
	markedListStyle.Width(w)
	markedListStyle.Height(height)
	m.Marked.SetSize(w, height)
}

func (m Model) View() (s string) {
	downloadableView := downloadableListStyle.Render(m.Downloadable.View())
	markedView := markedListStyle.Render(m.Marked.View())

	s = lipgloss.JoinHorizontal(
		lipgloss.Center,
		downloadableView,
		markedView,
	)

	s = lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		s,
	)
	return
}

func (m Model) Init() tea.Cmd { return nil }
