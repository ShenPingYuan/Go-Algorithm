package main

func main() {

}

func subarraySum(nums []int, k int) int {
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	count := 0
	m := map[int]int{}
	for i := 0; i < len(preSum); i++ {
		sub := preSum[i] - k
		if v, f := m[sub]; f {
			count += v
		}
		m[preSum[i]]++
	}

	return count
}
