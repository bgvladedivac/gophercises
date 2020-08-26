package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"flag"
)

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

func main() {
	ptrToFileName := flag.String("name of csv file", "problems.csv", "The name of the csv file - default to 'problems.csv'")
	records, err := getFileContent(*ptrToFileName)

	if err != nil {
		exit("Reader problem.")
	}


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
