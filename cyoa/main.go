package main

import (
	"cyoa/handlers"
	"log"
	"net/http"
)

var dataFileLocation, mainTemplateLocation string = "./data/gopher.json", "./templates/template.html"

func main() {
	stories := handlers.UnmarshallOverStoryStructHandler(dataFileLocation)
	handler := handlers.CustomHttpHandler{
		Stories:          stories,
		TemplateLocation: mainTemplateLocation,
	}
	// the port is hardcoded by purpose
	// usually 8082 would be free, unlike 8080
	log.Fatal(http.ListenAndServe(":8082", handler))
}
