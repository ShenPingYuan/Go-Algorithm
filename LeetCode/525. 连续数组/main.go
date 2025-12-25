package main

func main() {
	//nums = [0,1,1,1,1,1,0,0,0]
	nums := []int{0, 1, 1, 1, 1, 1, 0, 0, 0}
	findMaxLength(nums)
}

func findMaxLength(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	preSums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if v == 0 {
			v = -1
		}
		preSums[i+1] = preSums[i] + v
	}

	m := make(map[int]int)
	for i := 1; i < len(preSums); i++ {
		m[preSums[i]] = i
	}

	max := 0
	for i := 0; i < len(preSums); i++ {
		if v, f := m[preSums[i]]; f && v-i > max {
			max = v - i
		}
	}
	return max
}

func findMaxLength1(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); {
			sum += nums[j]
			j++
			if j == len(nums) {
				break
			}
			sum += nums[j]
			l := j - i + 1
			if sum == l/2 && l > maxLength {
				maxLength = l
			}
			j++
		}
		if len(nums)-i <= maxLength {
			return maxLength
		}
	}
	return maxLength
}
