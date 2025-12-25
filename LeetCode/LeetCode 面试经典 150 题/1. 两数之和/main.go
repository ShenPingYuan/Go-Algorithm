package main

import "fmt"

func main() {
	var nums []int = []int{2, 7, 11, 15}
	var target int = 13
	var result = twoSum(nums, target)

	fmt.Println(result)
}

//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
// 解法一
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return make([]int, 0)
}
