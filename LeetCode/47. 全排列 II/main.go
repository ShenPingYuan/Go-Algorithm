package main

import "fmt"

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	trace := []int{}
	visited := make([]bool, len(nums))
	selectSort(nums)
	var dfs func()
	dfs = func() {
		if len(trace) == len(nums) {
			res = append(res, append([]int(nil), trace...))
			return
		}
		levelSelected := map[int]bool{}
		for i := 0; i < len(nums); i++ {
			if levelSelected[nums[i]] {
				continue
			}
			if visited[i] {
				continue
			}
			levelSelected[nums[i]] = true
			visited[i] = true
			trace = append(trace, nums[i])
			dfs()
			trace = trace[:len(trace)-1]
			visited[i] = false
		}
	}
	dfs()
	return res
}
func selectSort(nums []int) {
	for s := 0; s < len(nums)-1; s++ {
		minIndex := s
		for i := s; i < len(nums); i++ {
			if nums[i] < nums[minIndex] {
				minIndex = i
			}
		}
		nums[s], nums[minIndex] = nums[minIndex], nums[s]
	}
}

func permuteUnique2(nums []int) [][]int {
	res := [][]int{}
	trace := []int{}
	visited := make([]bool, len(nums))
	selectSort(nums)
	var dfs func()
	dfs = func() {
		if len(trace) == len(nums) {
			res = append(res, append([]int{}, trace...))
			return
		}
		lastedNum := -11
		for i := 0; i < len(nums); i++ {
			if nums[i] == lastedNum {
				continue
			}
			if visited[i] {
				continue
			}
			lastedNum = nums[i]
			visited[i] = true
			trace = append(trace, nums[i])
			dfs()
			trace = trace[:len(trace)-1]
			visited[i] = false
		}
	}
	dfs()
	return res
}
