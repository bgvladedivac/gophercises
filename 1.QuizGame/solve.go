package main

import (
        "encoding/csv"
        "os"
        "fmt"
        "flag"
        "time"
)

const Ready = "yes"

var correctAnswers int

func getFileContent(filename string) (records[][] string, err error) {
        fd, err := os.Open(filename)

        if err != nil {
                exit("Problem with opening the file")
        }

        reader := csv.NewReader(fd)
        return reader.ReadAll()
}

func outputResult(totalQuestions, correctAnswers int) {
        fmt.Println("Number of total questions:", totalQuestions)
        fmt.Println("Number of corrected answers:", correctAnswers)
}

func exit(message string) {
        fmt.Println(message)
        os.Exit(3)
}

func provideQuestions(questions[][] string, done chan bool) {
        for index, value := range questions {
                question, answer := value[0], value[1]
                fmt.Println("Question N:=", index+1, "=>", question)

                var userAnswer string
                fmt.Scanf("%s", &userAnswer)

		if userAnswer == answer {
			correctAnswers ++
		}
        }

        done <- true
}


func main() {
        fileName := flag.String("name of csv file", "problems.csv", "The name of the csv file - default to 'problems.csv'")
        allowedTime := flag.Int("allowed time", 10, "The duration in seconds for which the quizz must be finished.")
	flag.Parse()

	var readyToStart string
	fmt.Println("Type 'yes' if you are ready to start")
	fmt.Scanf("%s", &readyToStart)

	if readyToStart != Ready {
		fmt.Println("Come back later if you are ready")
		os.Exit(0)
	}

        records, err := getFileContent(*fileName)


        if err != nil {
                exit("Reader problem.")
        }
	

        done := make(chan bool)
	timer := time.NewTimer(time.Duration(*allowedTime) * time.Second)
	
        go provideQuestions(records, done)

        select {
        case <- done:
                fmt.Println("Great job")
                fmt.Println("You finished with all questions.")
	case <- timer.C:
                fmt.Println("Time expired")
        }

	outputResult(len(records), correctAnswers)

}
