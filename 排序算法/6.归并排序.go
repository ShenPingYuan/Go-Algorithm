package main

import "fmt"

func main() {
	nums := []int{7, 2, 6, 4, 7, 9, 1, 17, 1, 0, 1, 5, 7, 8, 3, 2, 5, 1, 7, 9, 3, 1}
	mergeSort(nums)
	fmt.Println(nums)
}

func mergeSort(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, low int, high int) {
	if high <= low {
		return
	}
	mid := (high + low) / 2
	sort(nums, low, mid)
	sort(nums, mid+1, high)
	merge(nums, low, mid, high)
}

func merge(nums []int, low int, mid int, high int) {
	len := high - low + 1
	var temp []int = make([]int, len)
	lowIndex := low
	highIndex := mid + 1
	for i := 0; i < len; i++ {
		if lowIndex <= mid && highIndex <= high {
			if nums[lowIndex] <= nums[highIndex] {
				temp[i] = nums[lowIndex]
				lowIndex++
			} else {
				temp[i] = nums[highIndex]
				highIndex++
			}
		} else if lowIndex <= mid {
			temp[i] = nums[lowIndex]
			lowIndex++
		} else if highIndex <= high {
			temp[i] = nums[highIndex]
			highIndex++
		}
	}
	for i := 0; i < len; i++ {
		nums[low+i] = temp[i]
	}
}
