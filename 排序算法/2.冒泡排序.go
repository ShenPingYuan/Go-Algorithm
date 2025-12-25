package main

import "fmt"

func main() {
	var nums []int = []int{2, 4, 7, 1, 3, 6, 2, 7, 0, 6, 2, 1, 4, 6, 3}
	nums = bubbleSort(nums)

	fmt.Println(nums)
}

func bubbleSort(nums []int) []int {
	for sortedIndex := 0; sortedIndex < len(nums); sortedIndex++ {
		var swapped bool = false
		for j := len(nums) - 1; j > sortedIndex; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return nums
}
