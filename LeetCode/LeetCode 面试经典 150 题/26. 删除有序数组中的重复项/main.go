package main

func main() {
	nums := []int{1, 1, 2} // 输入数组

	removeDuplicates(nums) // 调用
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for i+1 < length && nums[i] == nums[i+1] {
			for k := i + 1; k+1 < length; k++ {
				nums[k] = nums[k+1]
			}
			length--
		}
	}
	return length
}

func removeDuplicates2(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	slow, fast := 0, 1
	for fast < length {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
