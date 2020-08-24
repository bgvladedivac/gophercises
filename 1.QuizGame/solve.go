package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"bufio"
)

func main() {
	fd, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println("Problem with opening the file")
	}

	reader := csv.NewReader(fd)

	records, err := reader.ReadAll()

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

		for scanner.Scan() {
			userGuess := scanner.Text()
			if userGuess == answer {
				counter += 1
			}
			break
		}

	}

	fmt.Println("Number of total questions:", totalQuestionsCounter)
	fmt.Println("Number of corrected answers:", counter)
}
