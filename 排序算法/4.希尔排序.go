package main

import "fmt"

func main() {
	nums := []int{4, 7, 3, 6, 2, 6, 5, 7, 3, 2, 1, 45, 34, 52, 346, 54, 63, 45, 4, 6, 357, 34, 5234, 525}
	shellSort(nums)
	fmt.Println(nums)
}

func shellSort(nums []int) {
	var n = len(nums)

	// 计算递增参数
	var d = 1
	for d < n/2 {
		d = d * 2
	}

	// 只要步长大于等于1，就执行插入排序，排序后步长减半
	for d >= 1 {
		// 改动二，sortedIndex 初始化为 h，而不是 1
		sortedIndex := d
		for sortedIndex < n {
			for i := sortedIndex; i >= d; i = i - d {
				if nums[i] >= nums[i-d] {
					break
				}
				nums[i], nums[i-d] = nums[i-d], nums[i]
			}
			sortedIndex++
		}
		d = d / 2
	}
}
