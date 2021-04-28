package gui

import "github.com/rivo/tview"

type Help struct {
	*tview.Table
}

var (
	helpHeaders = []string{"Key", "Description"}
	keybindings = []map[string]string{
		{"h": "show help (not work)"},
		{"q": "exit application (not work)"},
	}
)

func NewHelp() *Help {
	h := &Help{
		Table: tview.NewTable().SetBorders(true),
	}

	h.SetTitle("Help").
		SetTitleAlign(tview.AlignCenter)

	for col, v := range helpHeaders {
		h.SetCell(0, col, &tview.TableCell{
			Text:          v,
			NotSelectable: true,
			Align:         tview.AlignLeft,
		})
	}

	for i, elem := range keybindings {
		for key, desc := range elem {
			h.SetCell(i+1, 0, tview.NewTableCell(key))
			h.SetCell(i+1, 1, tview.NewTableCell(desc))
		}
	}

	return h
}
