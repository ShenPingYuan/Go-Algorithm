package main

import "fmt"

func main() {
	nums := [15]int{1, 2, 6, 2, 6, 2, 4, 6, 7, 9, 3, 2, 6, 8, 4}
	insertSort(nums[:])
	fmt.Println(nums)
}

func insertSort(nums []int) {
	for sortedIndex := 0; sortedIndex < len(nums); sortedIndex++ {
		for i := sortedIndex; i > 0; i-- {
			if nums[i] >= nums[i-1] {
				break
			}
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}
