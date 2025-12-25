package main

func main() {

}

func removeElement(nums []int, val int) int {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}
