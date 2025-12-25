package main

import "fmt"

func main() {
	fmt.Println("sort")
	nums := []int{3, 5, 1, 6, 7, 3, 8, 0, 2, 4, 5}
	fmt.Println(nums)
	sort(nums)
	fmt.Println(nums)
}

func sort(nums []int) {
	for sortedIndex := 0; sortedIndex < len(nums); sortedIndex++ {
		// sortedIndex往的后最小值的下标
		minIndex := sortedIndex
		// 从sortedIndex往后循环
		for j := sortedIndex; j < len(nums); j++ {
			// 如果下标j的元素小于minIndex位置的元素
			if nums[j] < nums[minIndex] {
				// 那么minIndex更新为j
				minIndex = j
			}
		}
		// 循环过后minIndex就是sortedIndex后面最小的元素的位置
		// 然后交换sortedIndex与minIndex位置的元素

		// old
		// var temp = nums[sortedIndex]
		// nums[sortedIndex] = nums[minIndex]
		// nums[minIndex] = temp

		// new
		nums[sortedIndex], nums[minIndex] = nums[minIndex], nums[sortedIndex]
	}
}
