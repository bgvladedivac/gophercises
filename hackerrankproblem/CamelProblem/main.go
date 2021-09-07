package main

import (
	"fmt"
	"unicode"
)

func getNumberOfWords(s string) int {
	result := 1
	for _, char := range s {
		if unicode.IsUpper(char) {
			result += 1
		}
	}
	return result
}

func main() {
	// test cases
	basicCase := "oneTwoThree"
	fmt.Println(getNumberOfWords(basicCase))
	basicCase = "saveChangesInTheEditor"
	fmt.Println(getNumberOfWords(basicCase))

	// input part
	var input string
	fmt.Scanf("%s", &input)
	numberOfWords := getNumberOfWords(input)
	fmt.Println(numberOfWords)
}
