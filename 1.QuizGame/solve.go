package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"bufio"
	"time"
)

func getFileContent(filename string) (records[][] string, err error) {
	fd, err := os.Open(filename)

	if err != nil {
		fmt.Println("Problem with opening the file.")
		os.Exit(-1)
	}

	reader := csv.NewReader(fd)
	return reader.ReadAll()
}

func provideUserSummary(totalQuestions, correctAnswers int) {
	fmt.Println("Number of total questions:", totalQuestions)
	fmt.Println("Number of corrected answers:", correctAnswers)
}

func isItDone(t Time)  {
	elapsedTime := time.Since(t)
	fmt.Println(elapsedTime)
}

func main() {
	start := time.Now()

	records, err := getFileContent("problems.csv")

	if err != nil {
		fmt.Println("Problem with the reader")
	}

	scanner := bufio.NewScanner(os.Stdin)

	counter := 0
	totalQuestionsCounter := 0

	for _, value := range records {
		totalQuestionsCounter += 1
		question := value[0]
		answer := value[1]
		fmt.Println("Question:", question)
		isItDone(start)
		for scanner.Scan() {
			userGuess := scanner.Text()
			if userGuess == answer {
				counter += 1
			}
			break
		}

	}

	provideUserSummary(totalQuestionsCounter, counter)

}
