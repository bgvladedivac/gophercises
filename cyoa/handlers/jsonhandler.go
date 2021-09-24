package handlers

import (
	"cyoa/entities"
	"encoding/json"
	"io/ioutil"
	"log"
)

func UnmarshallOverStoryStructHandler(gopherFileLocation string) entities.Story {
	content, err := ioutil.ReadFile(gopherFileLocation)
	if err != nil {
		log.Fatal("Error with opening the file", gopherFileLocation, "=> ", err)
	}

	var stories entities.Story

	err = json.Unmarshal([]byte(content), &stories)
	if err != nil {
		log.Fatal("Error while unmarshalling the json => ", err)
	}

	return stories
}
