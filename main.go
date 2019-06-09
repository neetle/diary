package main

import (
	"bytes"
	"diary/intent"
	"diary/view"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"strings"
	"text/template"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

func main() {
	// todo: derive from config
	state := &AppState{
		SavePath:  getDefaultPath(),
		Timestamp: time.Now().Unix(),
	}

	app := app.New()
	intents := &intent.Handler{
		App:      app,
		Delegate: state.delegate,
	}

	intents.Publish(intent.Start{})
	app.Run()
}

func getDefaultPath() string {
	user, err := user.Current()
	if err != nil {
		panic("couldn't get current user - " + err.Error())
	}

	return path.Join(
		user.HomeDir,
		"notes",
		"{{ .Timestamp }}.md",
	)
}

type AppState struct {
	SavePath  string
	Timestamp int64
}

func (a *AppState) delegate(app fyne.App, handler *intent.Handler, process intent.Intent) intent.Intent {

	switch p := process.(type) {
	case intent.Quit:
		app.Quit()
		return intent.None{}
	case intent.Start:
		view.SpawnMain(app, handler)
		return intent.None{}
	case intent.ContentUpdated:
		if len(p) >= 2 && p[len(p)-2:] == "\n\n" {
			return intent.ContentWrite(p)
		}
		return intent.None{}

	case intent.ContentWrite:
		return a.write(string(p))
	}

	panic(fmt.Sprintf("UNKNOWN INTENT %T", process))
}

func (a *AppState) write(content string) intent.Intent {
	t, err := template.New("filename").Parse(a.SavePath)
	checkThatWe(
		"parse the filename pattern supplied", err,
	)

	buf := &bytes.Buffer{}
	checkThatWe(
		"template the supplied filename pattern",
		t.Execute(buf, a),
	)

	dir := path.Dir(buf.String())

	checkThatWe(
		"ensure there's a parent folder for the note",
		os.MkdirAll(dir, 0700),
	)

	content = strings.TrimSpace(content) + "\n"
	checkThatWe(
		"write the file to the path ["+buf.String()+"]",
		ioutil.WriteFile(buf.String(), []byte(content), 0700),
	)

	return intent.Quit{}
}

func checkThatWe(while string, err error) {
	if err == nil {
		return
	}

	panicMsg := fmt.Sprintf("While trying to %s, the following error occured: %s", while, err.Error())
	panic(panicMsg)
}
