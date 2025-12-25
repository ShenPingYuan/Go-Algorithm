package main

import "fmt"

func main() {
	var nums []int = []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate1(nums []int) bool {
	var hashSet = make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if hashSet[nums[i]] {
			return true
		}
		hashSet[nums[i]] = true
	}
	return false
}

func containsDuplicate(nums []int) bool {
	var hashSet = make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, exists := hashSet[nums[i]]; exists {
			return true
		}
		hashSet[nums[i]] = struct{}{}
	}
	return false
}
