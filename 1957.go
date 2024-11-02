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
	var result strings.Builder
	result.Grow(len(s))

	var lastByte byte
	count := 0

	for i := 0; i < len(s); i++ {
		letter := s[i]

		if letter == lastByte {
			count++
		} else {
			count = 1
			lastByte = letter
		}

		if count <= 2 {
			result.WriteByte(letter)
		}
	}

	return result.String()
}
