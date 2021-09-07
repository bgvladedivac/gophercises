package main

import (
	"fmt"
	"net/http"
	"import from the other directory"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string {
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServerMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}


