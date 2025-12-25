package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums =[3,2,4]
	nums := []int{3, 2, 4}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	sort.Ints(nums2)
	left, right := 0, len(nums2)-1
	n1, n2 := -1, -1
	i, j := 0, 0
	for left < right {
		if nums2[left]+nums2[right] == target {
			i, j = nums2[left], nums2[right]
			break
		} else if nums2[left]+nums2[right] > target {
			right--
		} else if nums2[left]+nums2[right] < target {
			left++
		}
	}
	for k := 0; k < len(nums); k++ {
		if nums[k] == i && n1 == -1 {
			n1 = k

		} else if nums[k] == j && n2 == -1 {
			n2 = k
		}
	}
	if n1 < n2 {
		return []int{n1, n2}
	}
	return []int{n2, n1}
}

func twoSum1(nums []int, target int) []int {
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		left := target - nums[i]
		if v, found := hash[left]; found {
			return []int{v, i}
		}
		hash[nums[i]] = i
	}
	return nil
}
