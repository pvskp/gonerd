package adapters

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/pvskp/gonerd/cmd"
	"github.com/pvskp/gonerd/events"
	"github.com/pvskp/gonerd/structs"
)

func DownloadSingleFont(font list.Item, index int) events.FontDownloadedMsg {
	if item, ok := font.(structs.FontlistItem); ok {
		if err := cmd.DownloadFont(item.Content); err != nil {
			log.Println("Error while downloading font ", item.Content, " : ", err)
			return events.FontDownloadedMsg{
				Success: false,
				Index:   index,
			}
		}
	}
	return events.FontDownloadedMsg{
		Success: true,
		Index:   index,
	}
}

func DownloadFont(fonts []list.Item) events.FontDownloadedMsg {
	for i, v := range fonts {
		if item, ok := v.(structs.FontlistItem); ok {
			if err := cmd.DownloadFont(item.Content); err != nil {
				log.Println("Error while downloading font ", item, " : ", err)
				return events.FontDownloadedMsg{
					Success: false,
					Index:   i,
				}
			}
			return events.FontDownloadedMsg{
				Success: true,
				Index:   i,
			}
		}
	}

	return events.FontDownloadedMsg{
		Success: false,
		Index:   -1,
	}
}
