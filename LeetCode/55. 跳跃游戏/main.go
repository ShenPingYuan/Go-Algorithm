package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	pos := 0
	var dfs func() bool
	stepped := make([]bool, len(nums))
	dfs = func() bool {
		if pos >= len(nums)-1 {
			return true
		}
		length := nums[pos]
		for i := 1; i <= length; i++ {
			pos += i
			if pos >= len(nums)-1 {
				return true
			}
			if stepped[pos] {
				pos -= i
				continue
			}
			stepped[pos] = true
			if dfs() {
				return true
			}
			pos -= i
		}
		return false
	}
	return dfs()
}
func canJump2(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	currentStep := 0
	for currentStep+nums[currentStep] < len(nums)-1 {
		nextMaxStep := 0
		nextMaxStepIndex := 0
		for i := 1; i <= nums[currentStep]; i++ {
			if i+nums[currentStep+i] > nextMaxStep {
				nextMaxStep = i + nums[currentStep+i]
				nextMaxStepIndex = i
			}
		}
		if nextMaxStepIndex == 0 {
			return false
		}
		currentStep += nextMaxStepIndex
	}
	return true
}

func canJump0(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxStep := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]+i > maxStep {
			maxStep = nums[i] + i
		}
		if maxStep == i {
			return false
		}
		if maxStep >= len(nums)-1 {
			return true
		}
	}
	return true
}
