package main

import (
	"fmt"
	"strings"
)

func main() {
	result := makeFancyString("leeetcode")
	fmt.Println(result)
}

func makeFancyString(s string) string {
	runeSample := []rune(s)
	var result strings.Builder

	var lastRune rune
	count := 0

	for _, letter := range runeSample {
		if letter == lastRune {
			count++
		} else {
			count = 1
			lastRune = letter
		}

		if count <= 2 {
			result.WriteRune(letter)
		}
	}

	return result.String()
}
