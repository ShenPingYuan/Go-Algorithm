package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums = [-1,0,1,2,-1,-4]
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		// 跳过重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		f := nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			sum := f + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{f, nums[left], nums[right]})
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	res := [][]int{}
	for left < right {
		sum := nums[left] + nums[right]
		i, j := left, right
		if sum == target {
			res = append(res, []int{nums[left], nums[right]})
			for left < right && nums[j] == nums[right] {
				right--
			}

			for left < right && nums[i] == nums[left] {
				left++
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return res
}

func nSum(nums []int, start, n, target int) [][]int {
	l := len(nums)
	res := [][]int{}
	if n < 2 || n > l || start >= l {
		return res
	}
	if n == 2 {
		left, right := start, l-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	} else {
		for i := start; i < l-n; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			f := nums[i]
			nTarget := target - f
			n1Sums := nSum(nums, i+1, n-1, nTarget)
			for _, v := range n1Sums {
				res = append(res, append([]int{f}, v...))
			}
		}
	}
	return res
}
