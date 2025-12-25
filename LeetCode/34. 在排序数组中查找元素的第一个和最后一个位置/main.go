package main

func main() {

}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	leftBound := leftBound(nums, target)
	if leftBound == -1 {
		return []int{-1, -1}
	}
	rightBound := rightBound(nums, target)
	return []int{leftBound, rightBound}
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if left >= len(nums) {
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] <= target {
			left = mid + 1
		}
	}

	if right < 0 {
		return -1
	}
	if nums[right] == target {
		return right
	}
	return -1
}
