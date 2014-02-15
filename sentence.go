package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var punctuation map[string]bool = map[string]bool{
	".": true,
	"?": true,
	"!": true,
}
var markers map[string]string = map[string]string{
	"\"": "\"",
	"(":  ")",
	"[":  "]",
	"{":  "}",
}

func Sentence(input string) (sentences []string) {
	var current string
	currentMarker := ""

	for _, c := range input {
		current += string(c)

		_, isOpener := markers[string(c)]
		if isOpener && len(currentMarker) == 0 {
			currentMarker = string(c)
		} else if len(currentMarker) > 0 {
			closer := markers[currentMarker]
			if closer == string(c) {
				currentMarker = ""
			}
		}

		_, ok := punctuation[string(c)]
		if ok {
			// don't split if we're inside markers:
			if len(currentMarker) > 0 {
				continue
			}
			sentences = append(sentences, current)
			current = ""
		}
	}
	if len(current) > 0 {
		sentences = append(sentences, current)
	}
	return sentences
}

func main() {
	file := os.Stdin
	bufferedReader := bufio.NewReader(file)
	str, err := bufferedReader.ReadString('\n')
	if err != nil {
		fmt.Errorf("Error reading from stdin: %s", err)
	}
	sentences := Sentence(str)
	for _, s := range sentences {
		fmt.Println(strings.Trim(s, " "))
	}
}
