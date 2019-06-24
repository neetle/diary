package view

import (
	"diary/intent"

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

	w.SetContent(widget.NewVBox(
		m.TextEntry(),
	))

	w.Resize(size)
	w.CenterOnScreen()

	w.Show()

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
