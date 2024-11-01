package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	result := coinChange(coins, 11)
	fmt.Println(result)
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for a := 1; a <= amount; a++ {
		for c := 0; c < len(coins); c++ {
			if a-coins[c] >= 0 {
				dp[a] = min(dp[a], 1+dp[a-coins[c]])
			}
		}
	}
	if dp[amount] == (amount + 1) {
		return -1
	}
	return dp[amount]
}

func min(number1 int, number2 int) int {
	if number1 < number2 {
		return number1
	}
	return number2
}
