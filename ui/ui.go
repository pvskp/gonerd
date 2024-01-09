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
	docStyle = lipgloss.NewStyle().
		Margin(1, 1).
		Border(lipgloss.NormalBorder())
)

// var notification Notification = NewNotification()

const (
	DownloadableWidthPercentageSize = 0.6
	MarkedWidthPercentageSize       = 0.4
	NotificationWidthPercentageSize = 0.4

	DownloadableHeightPercentageSize = 1
	MarkedHeightPercentageSize       = 0.8
	NotificationHeightPercentageSize = 0.2

	stateDownloadable = iota
	stateMarked
	totalStates
)

type Model struct {
	Marked       structs.FontList
	Downloadable structs.FontList
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
		Marked:       structs.NewMarkedModel(mr),
		Downloadable: structs.NewDownloadableModel(d),
		state:        stateDownloadable,
	}

	log.Printf("[]list.Items = %v\n", m.Downloadable.Items())

	for _, v := range m.Downloadable.Items() {
		if item, ok := v.(structs.FontlistItem); ok {
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

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.width = msg.Width
		m.height = msg.Height

		m.Downloadable.Resize(
			int(float64(msg.Width-h)*DownloadableWidthPercentageSize),
			msg.Height-v,
		)

		m.Marked.Resize(
			int(float64(msg.Width-h)*MarkedWidthPercentageSize),
			msg.Height-v,
		)

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

// func setNotificationSize(n *Notification, width, height int) {
// 	w := int(float64(width) * NotificationWidthPercentageSize)
// 	h := int(float64(height) * NotificationHeightPercentageSize)
// 	n.Resize(w, h)
// }

func (m Model) View() (s string) {
	downloadableView := m.Downloadable.Render()
	markedView := m.Marked.Render()

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
