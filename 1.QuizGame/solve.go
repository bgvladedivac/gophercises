package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"flag"
	"time"
)

const Ready = "yes"

func getFileContent(filename string) (records[][] string, err error) {
	fd, err := os.Open(filename)

	if err != nil {
		exit("Problem with opening the file")
	}

	reader := csv.NewReader(fd)
	return reader.ReadAll()
}

func provideUserSummary(totalQuestions, correctAnswers int) {
	fmt.Println("Number of total questions:", totalQuestions)
	fmt.Println("Number of corrected answers:", correctAnswers)
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(3)
}

func provideQuestions(questions[][] string, done chan bool) {
	for index, value := range questions {
		question, _:= value[0], value[1]
		fmt.Println("Question N:=", index+1, "=>", question)

		var userAnswer string
		fmt.Scanf("%s", &userAnswer)
	}

	done <- true
}

func watchOutForTimeExpiration() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
}

func main() {
	fileName := flag.String("name of csv file", "problems.csv", "The name of the csv file - default to 'problems.csv'")
	//allowedTime := flag.Int64("allowed time", 10, "The duration in seconds for which the quizz must be finished.")
	records, err := getFileContent(*fileName)
	if err != nil {
		exit("Reader problem.")
	}

	done := make(chan bool)
	expiredTime := make(chan time.Time)

	go provideQuestions(records, done)
	go watchOutForTimeExpiration()

	select {
	case <- done:
		fmt.Println("Great job")
		fmt.Println("You finished with all questions.")
	case <- expiredTime:
		fmt.Println("Time expired")
	}

}
