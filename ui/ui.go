package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// titleStyle        = lipgloss.NewStyle().MarginLeft(2)

	docStyle = lipgloss.NewStyle().
			Margin(1, 2).
			Border(lipgloss.NormalBorder())

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("170"))

	// paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	// helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	// quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

const (
	stateDownloadable = iota
	stateMarked
	totalStates
)

type item string

func (i item) FilterValue() string {
	return string(i)
}

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type Model struct {
	Marked       list.Model
	Downloadable list.Model
	state        int
	width        int
	height       int
}

func NewModel() (m *Model) {
	d := []list.Item{
		item("Fira Code"),
		item("Anonymous Pro"),
		item("Inconsolata"),
	}

	mr := []list.Item{
		item("Hack"),
		item("Hack"),
		item("Hack"),
	}

	m = &Model{
		Marked:       list.New(mr, itemDelegate{}, 0, 0),
		Downloadable: list.New(d, itemDelegate{}, 0, 0),
		state:        stateDownloadable,
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
				m.Marked.InsertItem(len(m.Marked.Items()), m.Downloadable.SelectedItem())
			}

		case stateMarked:
			switch msg.String() {
			case "d":
				m.Marked.RemoveItem(m.Downloadable.Index())
			}
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.width = msg.Width
		m.height = msg.Height
		m.Downloadable.SetSize(msg.Width-h, msg.Height-v)
		m.Marked.SetSize(msg.Width-h, msg.Height-v)
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

func (m Model) View() string {
	downloadable := m.Downloadable
	marked := m.Marked
	s := lipgloss.JoinHorizontal(lipgloss.Center, downloadable.View(), marked.View())
	s = lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		s,
	)
	return s

}

func (m Model) Init() tea.Cmd { return nil }
