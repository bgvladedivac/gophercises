package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"bufio"
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

	scanner := bufio.NewScanner(os.Stdin)

	counter := 0

	for _, value := range records {
		question, answer  := value[0], value[1]
		fmt.Println("Question:", question)

		for scanner.Scan() {
			userGuess := scanner.Text()
			if userGuess == answer {
				counter += 1
			}

			break
		}

	}

	provideUserSummary(len(records), counter)

}
