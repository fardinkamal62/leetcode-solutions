package main

import (
	"fmt"
	"strings"
)

func main() {
	result := compressedString("cccccccccc")
	fmt.Println(result)
}

// Solution 2
// 93ms, 8.7mb
func compressedString(s string) string {
	var resultString strings.Builder
	key := string(s[0])
	count := 0

	for i := 0; i < len(s); i++ {
		letter := string(s[i])
		if letter == key {
			count++
		} else {
			if count > 9 {
				resultString.WriteString(fmt.Sprintf("9%s", key))
				count -= 9
				i--
			} else {
				resultString.WriteString(fmt.Sprintf("%d%s", count, key))
				key = letter
				count = 1
			}
		}
	}

	if count > 9 {
		for count > 0 {
			if count > 9 {
				resultString.WriteString(fmt.Sprintf("9%s", key))
				count -= 9
			} else {
				resultString.WriteString(fmt.Sprintf("%d%s", count, key))
				count = 0
			}
		}
	} else {
		resultString.WriteString(fmt.Sprintf("%d%s", count, key))
	}
	return resultString.String()
}

// Solution 1
/*
42ms, 9mb

func compressedString(s string) string {
	var resultString strings.Builder

	for i := 0; i < len(s); i++ {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			count++
			i++
		}
		if count > 9 {
			resultString.WriteString(fmt.Sprintf("%d%c", 9, s[i]))
			count -= 9
			for count > 9 {
				resultString.WriteString(fmt.Sprintf("%d%c", 9, s[i]))
				count -= 9
			}
			resultString.WriteString(fmt.Sprintf("%d%c", count, s[i]))
		} else {
			resultString.WriteString(fmt.Sprintf("%d%c", count, s[i]))
		}
	}

	return resultString.String()
}

*/
