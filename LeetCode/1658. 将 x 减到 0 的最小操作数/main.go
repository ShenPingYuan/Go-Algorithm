package main

import "fmt"

func main() {
	//nums = [1,1,4,2,3] x=5
	nums := []int{3, 2, 20, 1, 1, 3}
	x := 10
	result := minOperations(nums, x)
	fmt.Println(result)
}

func minOperations(nums []int, x int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	target := sum - x
	window := 0
	res := -1
	left, right := 0, 0
	for right < len(nums) {
		window += nums[right]
		right++

		for left < right && window > target {
			window -= nums[left]
			left++
		}

		if window == target {
			if right-left > res {
				res = right - left
			}
		}
	}
	if res == -1 {
		return -1
	}
	return len(nums) - res
}
