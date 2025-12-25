package main

import (
	"fmt"
)

func main() {
	res := permute([]int{1, 2, 4})
	fmt.Println(res)
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	trace := make([]int, 0, len(nums))
	selected := make([]bool, len(nums))
	fmt.Println("trace: ", len(trace), cap(trace))

	backTrace(&res, trace, nums, selected)
	return res
}

func backTrace(res *[][]int, trace []int, nums []int, selected []bool) {
	if len(trace) == len(nums) {
		temp := make([]int, len(trace))
		copy(temp, trace)
		*res = append(*res, temp)
		return
	}
	for i, v := range nums {
		if selected[i] {
			continue
		}
		trace = append(trace, v)
		selected[i] = true
		backTrace(res, trace, nums, selected)
		trace = trace[:len(trace)-1]
		selected[i] = false
	}
}
