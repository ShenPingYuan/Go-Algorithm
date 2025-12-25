package main

func main() {
	//nums = [10,5,2,6] k = 100
	nums := []int{10, 5, 2, 6}
	k := 100
	numSubarrayProductLessThanK(nums, k)
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	left := 0
	res := 0
	for left < len(nums) {
		window := 1
		right := left
		window *= nums[right]
		for right < len(nums) && window < k {
			res++
			right++
			window *= nums[right]
		}
		window /= nums[left]
		left++
	}
	return res
}

func numSubarrayProductLessThanK1(nums []int, k int) int {
	left, right := 0, 0
	res := 0
	window := 1
	for right < len(nums) {
		window *= nums[right]
		right++

		for left < right && window >= k {
			window /= nums[left]
			left++
		}

		res += right - left
	}
	return res
}
