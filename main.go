package main

import (
	"bytes"
	"diary/intent"
	"diary/view"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"strings"
	"text/template"
	"time"
)

func main() {
	file, err := os.OpenFile(getDefaultLogPath(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("starting up")

	defer func() {
		log.Println("finished")
		if e := recover(); e != nil {
			log.Println(e)
		}
	}()

	// todo: derive from config
	state := &AppState{
		SavePath:  getDefaultPath(),
		Timestamp: time.Now().Unix(),
		Intents:   intent.NewHandler(),
	}

	app := view.SpawnMain(state.Intents)

	app.Run(func() {
		select {
		case <-state.Intents.Quit:
		case content := <-state.Intents.ContentWrite:
			state.write(content)
		}
	
		app.Quit()
	})
}

func getDefaultLogPath() string {
	user, err := user.Current()
	if err != nil {
		panic("couldn't get current user - " + err.Error())
	}

	return path.Join(
		user.HomeDir,
		"notes",
		"diary.log",
	)
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
	Intents   *intent.Handler
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
