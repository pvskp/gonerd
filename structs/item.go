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

type FontlistItem struct {
	Content           string
	itemStyle         lipgloss.Style
	selectedItemStyle lipgloss.Style
}

func (i FontlistItem) FilterValue() string {
	return i.Content
}

func (i FontlistItem) Title() string       { return i.Content }
func (i FontlistItem) Description() string { return "" }

func NewItem(content string) FontlistItem {
	return FontlistItem{
		Content:           content,
		itemStyle:         itemStyle,
		selectedItemStyle: selectedItemStyle,
	}
}
