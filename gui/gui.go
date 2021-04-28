package gui

import "github.com/rivo/tview"

type Panel int

type Gui struct {
	CurrentPanel Panel
	Config       Config
	App          *tview.Application
}

// Create new gui
func New(config Config) *Gui {
	gui := &Gui{
		Config: config,
		App:    tview.NewApplication(),
	}
	return gui
}

func (gui *Gui) Run() error {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	contactform := NewContact()

	grid := tview.NewGrid().
		SetRows(0, 4, 5).
		SetBorders(true).
		AddItem(newPrimitive("logged callsign list"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("result of searching logs for input callsign"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(contactform, 1, 0, 1, 3, 0, 0, false)

	if err := gui.App.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		return err
	}
	return nil
}
