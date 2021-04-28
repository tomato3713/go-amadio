package gui

type Panel int

type Gui struct {
	CurrentPanel Panel
	Config       Config
}

// Create new gui
func New(config Config) *Gui {
	gui := &Gui{
		Config: config,
	}
	return gui
}

func (gui *Gui) Run() error {
	return nil
}
