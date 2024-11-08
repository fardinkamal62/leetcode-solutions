package main

import (
	"fmt"
)

func main() {
	result := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, exists := hashMap[complement]; exists {
			return []int{index, i}
		}
		hashMap[num] = i
	}

	return []int{}
}
