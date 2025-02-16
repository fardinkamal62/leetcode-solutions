package main

import "fmt"

func strStr(haystack string, needle string) int {
	var result int = -1

	n0 := needle[0]

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == n0 {
			if len(haystack)-i < len(needle) {
				break
			}
			if haystack[i:i+len(needle)] == needle {
				result = i
				break
			}
		}
	}

	return result
}

func main() {
	haystack := "hello"
	needle := "ll"
	result := strStr(haystack, needle)
	fmt.Println(result)
}
