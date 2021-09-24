package tests

import (
	"cyoa/handlers"
	"fmt"
	"testing"
)

const dataFileLocation string = "../data/gopher.json"

func TestUnmarshallOverStoryStruct(t *testing.T) {
	t.Parallel()
	stories := handlers.UnmarshallOverStoryStructHandler(dataFileLocation)

	for key, _ := range stories {
		fmt.Println(key)
	}

}
