package main

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	slow, fast := 0, 1
	dup := 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			dup = 1
			slow++
			nums[slow] = nums[fast]
		} else {
			dup++
			if dup < 3 {
				slow++
				nums[slow] = nums[fast]
			}
		}
		fast++
	}
	return slow + 1
}
