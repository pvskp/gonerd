package structs

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FontList struct {
	list       list.Model
	modelStyle lipgloss.Style
	titleStyle lipgloss.Style
	width      int
	height     int
}

func (d *FontList) Update(msg tea.Msg) (FontList, tea.Cmd) {
	l, cmd := d.list.Update(msg)
	d.list = l

	return *d, cmd
}

func (d FontList) Items() (l []list.Item) {
	l = d.list.Items()
	return
}

func (d FontList) InsertItem(index int, item list.Item) tea.Cmd {
	return d.list.InsertItem(index, item)
}

func (d FontList) RemoveItem(index int) {
	d.list.RemoveItem(index)
}

func (d FontList) SelectedItem() list.Item {
	return d.list.SelectedItem()
}

func (d FontList) Index() int {
	return d.list.Index()
}

func (d FontList) Render() string {
	return d.modelStyle.Render(d.View())
}

func (d FontList) View() (s string) {
	return d.list.View()
}

func (d *FontList) Resize(width, height int) {
	d.width = width
	d.height = height
	d.modelStyle = d.modelStyle.Width(width)
	d.modelStyle = d.modelStyle.Height(height)
	d.list.SetSize(width, height)
}

func NewDownloadableModel(itemList []list.Item) FontList {
	l := list.New(itemList, ItemDelegate{}, 0, 0)
	l.Title = "Downloadable"
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)

	return FontList{
		list: l,

		modelStyle: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true, true, true, true).
			Align(lipgloss.Left),

		titleStyle: lipgloss.NewStyle().
			Background(lipgloss.Color("205")).
			Align(lipgloss.Center),

		width:  0,
		height: 0,
	}
}

func NewMarkedModel(itemList []list.Item) FontList {
	l := list.New(itemList, ItemDelegate{}, 0, 0)
	l.Title = "Marked"
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)

	return FontList{
		list: l,

		modelStyle: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Align(lipgloss.Left),

		titleStyle: lipgloss.NewStyle().
			Background(lipgloss.Color("205")).
			Align(lipgloss.Center),

		width:  0,
		height: 0,
	}
}
