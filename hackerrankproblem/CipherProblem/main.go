package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getShiftedRune(capitalRune bool, char rune, shifting int) string {
	if capitalRune == true {
		char = unicode.ToLower(char)
	}
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	charIndex := strings.Index(alphabet, string(char))
	if charIndex == len(alphabet)-1 {
		charIndex = -1
	}
	startIndex := charIndex + 1
	counter := 1

	for {
		if counter == shifting {
			if capitalRune == true {
				return string(unicode.ToUpper(rune(alphabet[startIndex])))
			} else {
				return string(alphabet[startIndex])
			}

		}

		counter++
		startIndex++
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inputCounter := 0
	numberOfLines := 2
	var inputs [2]string
	for {
		if inputCounter == numberOfLines {
			break
		}

		scanner.Scan()
		inputs[inputCounter] = scanner.Text()
		inputCounter++
	}
	fmt.Println(inputs)
	basicInput := inputs[0]
	shifting, _ := strconv.Atoi(inputs[1])

	result := ""
	for _, char := range basicInput {
		if unicode.IsLetter(char) && unicode.IsUpper(char) {
			result += getShiftedRune(true, char, shifting)
		} else if unicode.IsLetter(char) {
			result += getShiftedRune(false, char, shifting)
		} else {
			result += string(char)
		}
	}
	fmt.Println(result)
}
