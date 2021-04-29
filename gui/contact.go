package gui

import "github.com/rivo/tview"

type ContactForm struct {
	Form   *tview.Flex
	Qso    *tview.Form
	Clock  *tview.TextView
	Common *tview.Form
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

func NewContact() *ContactForm {
	qso := tview.NewForm().SetHorizontal(true)
	common := tview.NewForm().SetHorizontal(true)
	clock := tview.NewTextView()

	qso.AddInputField("Call", "", 10, nil, nil).
		AddInputField("Snt", "", 4, nil, nil).
		AddInputField("Rst", "", 4, nil, nil).
		AddInputField("Comment", "", 30, nil, nil)
	common.AddDropDown("freq", freqList, 0, nil).
		AddDropDown("mode", modeList, 0, nil).
		SetBorderPadding(0, 0, 1, 1)

	clock.SetMaxLines(3).
		SetScrollable(false).
		SetText("YYYY/MM/DD 00:00:00")

	form := tview.NewFlex()
	form.SetDirection(tview.FlexRow).
		AddItem(qso, 3, 1, true).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(common, 30, 1, true).
				AddItem(clock, 0, 1, false),
			0, 1, true)

	contact := &ContactForm{
		Form:   form,
		Qso:    qso,
		Clock:  clock,
		Common: common,
	}

	return contact
}
