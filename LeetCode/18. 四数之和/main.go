package main

import "fmt"

func main() {
	// nums =[2,2,2,2,2]
	nums := []int{0, 0, 0, 0}
	target := 0
	result := nSum(nums, 0, 4, target)
	fmt.Println(result)
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
		for i := start; i < l-n+1; i++ {
			if i > start && nums[i] == nums[i-1] {
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
