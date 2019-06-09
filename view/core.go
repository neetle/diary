package view

import (
	"diary/intent"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type Main struct {
	out *intent.Handler
	app fyne.App
}

var size = fyne.NewSize(600, 80)

func SpawnMain(app fyne.App, out *intent.Handler) {
	m := &Main{
		out: out,
		app: app,
	}

	//todo - set width and height to be not gross
	w := m.app.NewWindow("Diary")
	w.RequestFocus()

	w.SetContent(widget.NewVBox(
		m.TextEntry(),
	))

	w.Resize(size)
	w.CenterOnScreen()

	w.Show()

}

func (m *Main) Submit() (string, func()) {
	return "Submit", func() {
		m.out.Publish(intent.Quit{})
	}
}

func (m *Main) TextEntry() *widget.Entry {
	box := widget.NewMultiLineEntry()
	box.OnChanged = func(content string) {
		m.out.Publish(intent.ContentUpdated(content))
	}
	box.Resize(size)
	box.FocusGained()

	return box
}
