package main

func main() {

}

func coinChange(coins []int, amount int) int {
	catch := make([]int, amount+1)
	return dp(coins, amount, catch)
}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func dp(coins []int, amount int, catch []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	if catch[amount] != 0 {
		return catch[amount]
	}

	minSubRes := 1<<31 - 1
	for _, v := range coins {
		var subRes int = dp(coins, amount-v, catch)
		if subRes == -1 {
			continue
		}
		minSubRes = min(minSubRes, subRes)
	}
	if minSubRes == 1<<31-1 {
		catch[amount] = -1
		return -1
	}
	catch[amount] = minSubRes + 1
	return minSubRes + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
