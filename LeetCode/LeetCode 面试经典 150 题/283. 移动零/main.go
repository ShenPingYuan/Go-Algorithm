package main

import "fmt"

func main() {
	nums := []int{1, 0, 1}
	slice := make([]int, 2, len(nums))
	copy(slice, nums)
	fmt.Println("slice:", slice)
	//moveZeroes(nums)
}

func moveZeroes(nums []int) {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[left] != 0 {
			left++
			continue
		}
		if nums[left] == 0 && nums[i] != 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
}

func moveZeroes1(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}
