package main

func main() {

}

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums)+1)
	prefix[0] = 1
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = nums[i] * prefix[i]
	}

	suffix := make([]int, len(nums)+1)
	suffix[len(nums)] = 1
	for i := len(nums) - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = prefix[i] * suffix[i+1]
	}
	return res
}
