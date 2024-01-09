package structs

import "github.com/charmbracelet/lipgloss"

type Notification struct {
	width       int
	height      int
	windowStyle lipgloss.Style
}

func (n *Notification) Resize(width, height int) {
	n.windowStyle.Width(width)
	n.windowStyle.Height(height)
}

func (n *Notification) View() (v string) { return }
func NewNotification() (n *Notification) {
	n = &Notification{
		windowStyle: lipgloss.NewStyle().
			Background(lipgloss.Color("205")).
			Border(lipgloss.NormalBorder()),

		width:  0,
		height: 0,
	}
	return
}
