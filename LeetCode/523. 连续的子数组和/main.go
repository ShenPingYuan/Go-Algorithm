package main

func main() {
	// nums = [23,2,4,6,6] k = 7
	nums := []int{23, 2, 4, 6, 6}
	k := 7
	checkSubarraySum(nums, k)
}

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	preSums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSums[i+1] = (preSums[i] + nums[i]) % k

	}
	m := map[int]int{}
	for i := 0; i < len(preSums); i++ {
		if v, f := m[preSums[i]]; f {
			if i-v >= 2 {
				return true
			}
		} else {
			m[preSums[i]] = i
		}
	}
	return false
}
