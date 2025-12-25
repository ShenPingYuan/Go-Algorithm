package main

func main() {

}

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}
