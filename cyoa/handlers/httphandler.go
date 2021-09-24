package handlers

import (
	"cyoa/entities"
	"html/template"
	"log"
	"net/http"
)

type CustomHttpHandler struct {
	Stories          entities.Story
	TemplateLocation string
}

func (ch CustomHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	storyFound := false
	for story, _ := range ch.Stories {
		if story == path {
			storyFound = true
		}
	}

	if !storyFound {
		missingStoryInfo := "Sorry, the story is not found"
		w.Write([]byte(missingStoryInfo))
		return
	}

	tmpl, err := template.ParseFiles(ch.TemplateLocation)

	if err != nil {
		log.Fatal("Problem with parsing the template")
	}

	tmpl.Execute(w, ch.Stories[path])
	return
}
