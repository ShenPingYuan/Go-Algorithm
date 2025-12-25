package main

func main() {
}

func sortColors2(nums []int) {
	if len(nums) <= 1 {
		return
	}
	left, right := 0, len(nums)-1
	m := 0
	for m <= right {
		if nums[m] == 0 {
			nums[left], nums[m] = nums[m], nums[left]
			left++
		} else if nums[m] == 2 {
			nums[right], nums[m] = nums[m], nums[right]
			right--
		} else {
			m++
		}
		if m < left {
			m = left
		}
	}
}
