package view

import (
	"diary/intent"

	"diary/hotkey"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type App interface {
	Run(do func())
	Quit()
}

type Main struct {
	out *intent.Handler
	app fyne.App
}

var size = fyne.NewSize(600, 80)

func SpawnMain(out *intent.Handler) App {
	m := &Main{
		out: out,
		app: app.New(),
	}

	//todo - set width and height to be not gross
	w := m.app.NewWindow("Diary")
	w.RequestFocus()

	textEntry := m.TextEntry()

	w.SetContent(widget.NewVBox(
		textEntry,
	))

	w.Canvas().Focus(textEntry)

	w.Resize(size)
	w.CenterOnScreen()

	w.Show()

	err := hotkey.Register(1, func() {
		w.RequestFocus()
	})

	if err != nil {
		textEntry.SetText(err.Error())
		textEntry.SetReadOnly(true)
	}

	return m
}

func (m *Main) Run(do func()) {
	go do()
	m.app.Run()
}

func (m *Main) Quit() {
	m.app.Quit()
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
