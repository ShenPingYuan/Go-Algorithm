package main

func main() {

}

func longestOnes(nums []int, k int) int {
	window := 0
	windowZero := 0
	left, right := 0, 0
	res := 0
	for right < len(nums) {
		new := nums[right]
		window++
		if new == 0 {
			windowZero++
		}

		right++
		for left < right && windowZero > k {
			d := nums[left]
			if d == 0 {
				windowZero--
			}
			window--
			left++
		}
		if window > res {
			res = window
		}
	}
	return res
}
