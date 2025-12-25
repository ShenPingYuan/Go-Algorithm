package main

func main() {

}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	n := len(nums) - 1
	for left <= right {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			res[n] = nums[left] * nums[left]
			left++
		} else {
			res[n] = nums[right] * nums[right]
			right--
		}
		n--
	}
	return res
}
