package main

import (
	"slices"
)

func main() {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
}

func reverseString(s []byte) {
	slices.Reverse(s)
}
