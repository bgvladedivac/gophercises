package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func handleError(message string, e error) {
	fmt.Println(message)
	log.Fatal(e)
}

type Link struct {
	Href string
	Text string
}

func CreateNewLink(href, text string) *Link {
	return &Link{href, text}
}

func (link *Link) String() string {
	return fmt.Sprintf("Link{Href:%s, Text:%s,}", link.Href, link.Text)
}

func printLinks(links []*Link) {
	for _, value := range links {
		fmt.Println(value)
	}
}

func findIndexWhereOpeningAnchorTagFinishes(pattern, text string) int {
	result := -1
	closingBracket := ">"
	patternIndex := strings.Index(text, pattern)

	if patternIndex == -1 {
		return result
	}

	indexForSearching := patternIndex + len(pattern) - 1
	for {
		indexForSearching += 1
		if strings.Compare(string(text[indexForSearching]), closingBracket) == 0 {
			return indexForSearching
		}
	}

}

func extractTextAfterIndexUntilFoundClosingAnchorTag(index int, text string) string {
	result := ""
	for {
		if strings.Compare(string(text[index+1]), "<") == 0 {
			break
		}

		result = result + string(text[index])
		index += 1
	}
	return result
}

func main() {
	files := []string{"./ex1.html", "./ex2.html", "./ex3.html", "./ex4.html"}

	for _, file := range files {
		fd, err := os.Open(file)

		if err != nil {
			handleError("Problem with opening the file", err)
		}
		fmt.Println("Itearting over file", fd.Name())

		fileContent, err := ioutil.ReadFile(file)

		if err != nil {
			handleError("Problem with reading the content of the file", err)
		}

		pointerToTokenizer := html.NewTokenizer(fd)
		errorTokenFound := false

		links := make([]*Link, 0)
		for !errorTokenFound {
			tokenType := pointerToTokenizer.Next()
			switch {
			case tokenType == html.ErrorToken:
				errorTokenFound = true
			case tokenType == html.StartTagToken:
				token := pointerToTokenizer.Token()
				if token.Data == "a" {
					for _, value := range token.Attr {
						if value.Key == "href" {
							afterPattern := value.Val + "\""

							index := findIndexWhereOpeningAnchorTagFinishes(afterPattern, string(fileContent))
							text := extractTextAfterIndexUntilFoundClosingAnchorTag(index, string(fileContent))
							text = strings.TrimSpace(text[1:])
							link := CreateNewLink(value.Val, text)
							links = append(links, link)
						}

					}
				}
			}
		}

		for _, value := range links {
			fmt.Println(value)
		}
	}

}
