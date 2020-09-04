package urlshort

import (
	"net/http"
)


func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandleFunc {
	//
	return nil
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandleFunc, error) {
	// 
	return nil, nil
}
