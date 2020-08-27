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

func inScopeWithTime(seconds int64) {
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("Your time has ended.")
	fmt.Println("You had", seconds, "seconds.")
	os.Exit(1)
}

func promptUser(allowedTime int64) {
	//if ready start out the goutine to check if time has been exceeded
	fmt.Println("Type 'yes' if you are ready to start")
	var ready string
	fmt.Scanf("%s", &ready)

	if ready == Ready {
		go inScopeWithTime(allowedTime)
	} else {
		promptUser(allowedTime)
	}
}

func main() {
	fileName := flag.String("name of csv file", "problems.csv", "The name of the csv file - default to 'problems.csv'")
	allowedTime := flag.Int64("allowed time", 30, "The duration in seconds for which the quizz must be finished.")

	records, err := getFileContent(*fileName)

	if err != nil {
		exit("Reader problem.")
	}


	promptUser(*allowedTime)

	counter := 0
	for _, value := range records {
		question, answer  := value[0], value[1]
		fmt.Println("Question:", question)
		var userAnswer string
		fmt.Scanf("%s", &userAnswer)

		if userAnswer == answer {
			counter++
		}

	}

	provideUserSummary(len(records), counter)
}
