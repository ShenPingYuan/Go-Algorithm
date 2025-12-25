package main

func main() {

}

func pivotIndex(nums []int) int {
	l := len(nums)
	sums := make([]int, l+1)
	for i := range l {
		sums[i+1] = nums[i] + sums[i]
	}

	for i := 0; i < l; i++ {
		if sums[i] == (sums[l] - sums[i+1]) {
			return i
		}
	}
	return -1
}
