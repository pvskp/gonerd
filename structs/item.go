package structs

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	itemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(1).
				Foreground(lipgloss.Color("175"))
)

type Item struct {
	Content           string
	itemStyle         lipgloss.Style
	selectedItemStyle lipgloss.Style
}

func (i Item) FilterValue() string {
	return i.Content
}

func (i Item) Title() string       { return i.Content }
func (i Item) Description() string { return "" }

func NewItem(content string) Item {
	return Item{
		Content:           content,
		itemStyle:         itemStyle,
		selectedItemStyle: selectedItemStyle,
	}
}
