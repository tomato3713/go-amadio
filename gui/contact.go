package gui

import "github.com/rivo/tview"

type ContactForm struct {
	*tview.Form
}

var (
	freqList = []string{
		"135 kHz",                                                                       // LF
		"1.9 MHz",                                                                       // MF
		"3.5 MHz", "3.8 MHz", "7 MHz", "10 MHz", "14 MHz", "21 MHz", "24 MHz", "28 MHz", // HF
		"50 MHz", "144 MHz", // VHF
		"430 MHz", "1.2 GHz", "2400 GHz", // UHF
		"5.6 GHz", // SHF
	}

	modeList = []string{
		"SSB", "FM", "AM", "CW", "RTTY",
	}
)

func NewContact() *tview.Form {
	form := tview.NewForm().
		SetHorizontal(true).
		AddInputField("Call", "", 10, nil, nil).
		AddInputField("Snt", "", 4, nil, nil).
		AddInputField("Rst", "", 4, nil, nil).
		AddInputField("Comment", "", 20, nil, nil)

	form.
		AddDropDown("freq", freqList, 0, nil).
		AddDropDown("mode", modeList, 0, nil)

	form.SetTitle("contact form").SetTitleAlign(tview.AlignLeft)

	return form
}
