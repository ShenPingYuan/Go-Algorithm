package main

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			fast++
			continue
		}
		slow++
		nums[slow] = nums[fast]
		fast++
	}
	return slow + 1
}
