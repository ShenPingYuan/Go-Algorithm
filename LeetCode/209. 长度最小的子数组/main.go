package main

import "fmt"

func main() {
	// var min1 int32 = 1 << 31-1
	// fmt.Println(min1) // 输出：-2147483648
	var min2 int32 = -1 << 31
	fmt.Println(min2) // 输出：-2147483648

	var max int32 = 1<<31 - 1
	max2 := max + 1
	fmt.Println(max) // 输出：2147483647
	fmt.Println(max2)
}

func minSubArrayLen(target int, nums []int) int {
	window := 0
	left, right := 0, 0
	min := 1<<31 - 1
	for right < len(nums) {
		new := nums[right]
		window += new
		right++

		for window >= target && left < right {
			if right-left < min {
				min = right - left
			}
			d := nums[left]
			window -= d
			left++
		}
	}
	if min == 1<<31-1 {
		return 0
	}
	return min
}
