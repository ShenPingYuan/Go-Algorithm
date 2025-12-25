package main

func main() {

}

func subarraysDivByK(nums []int, k int) int {
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	count := 0
	m := map[int]int{}
	for i := 0; i < len(preSum); i++ {
		mod := preSum[i] % k
		if mod < 0 {
			mod += k
		}
		if v, f := m[mod]; f {
			count += v
		}
		m[mod]++
	}
	return count
}
