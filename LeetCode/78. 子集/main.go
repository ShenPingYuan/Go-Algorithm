package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	trace := []int{}
	backtrace(&res, trace, nums, 0)
	return res
}

func backtrace(res *[][]int, trace []int, nums []int, index int) {
	co := make([]int, len(trace))
	copy(co, trace)
	*res = append(*res, co)
	for i := index; i < len(nums); i++ {
		trace = append(trace, nums[i])
		backtrace(res, trace, nums, i+1)
		trace = trace[:len(trace)-1]
	}
}
