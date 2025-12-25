package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{4, 7, 3, 6, 2, 6, 5, 7, 3, 2, 1, 45, 34, 52, 346, 54, 63, 45, 4, 6, 357, 34, 5234, 525}
	sort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func sort(nums []int, start int, end int) {
	if start >= end {
		return
	}
	var pos = partition2(nums, start, end)
	sort(nums, start, pos-1)
	sort(nums, pos+1, end)
}

func partition2(nums []int, start int, end int) int {
	var pivot = nums[start]
	var startIndex = start
	var endIndex = end
	var isLeft bool = false
	for startIndex < endIndex {
		if !isLeft {
			if nums[endIndex] >= pivot {
				endIndex--
			} else {
				nums[startIndex] = nums[endIndex]
				startIndex++
				isLeft = true
			}
		} else {
			if nums[startIndex] <= pivot {
				startIndex++
			} else {
				nums[endIndex] = nums[startIndex]
				endIndex--
				isLeft = false
			}
		}
	}
	if isLeft {
		nums[endIndex] = pivot
		return endIndex
	}
	nums[startIndex] = pivot
	return startIndex
}

func partition1(nums []int, start int, end int) int {
	// 第一个元素作为分割元素
	var pivot = nums[start]

	// 数组长度
	len := end - start + 1

	// 初始化一个临时数组
	tempNums := make([]int, len)

	startIndex := 0     // 临时数组的头部序号
	endIndex := len - 1 // 临时数据的尾部序号

	// 从第二个元素开始遍历nums
	for i := start + 1; i <= end; i++ {
		if nums[i] <= pivot { // 如果小于等于分割元素，放到临时数组头部
			tempNums[startIndex] = nums[i]
			startIndex++
		} else { // 如果大于分割元素，放到临时数组尾部
			tempNums[endIndex] = nums[i]
			endIndex--
		}
	}
	// 再将分割元素放到中间那个空位
	tempNums[startIndex] = pivot

	// 最后再将临时数组中的元素覆盖原来的nums
	for i := 0; i < len; i++ {
		nums[start+i] = tempNums[i]
	}

	// 返回分割元素pivot在原数组中的位置
	return start + startIndex
}

// 优化：
func quickSort(nums []int) {
	quickSortHelper(nums, 0, len(nums)-1)
}

func quickSortHelper(nums []int, left, right int) {
	// 小数组使用插入排序
	if right-left <= 10 {
		insertionSort(nums, left, right)
		return
	}

	if left < right {
		// 三数取中选择基准
		pivotIndex := medianOfThree(nums, left, right)
		// 将基准移到开始位置
		nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]

		// 分区
		p := partition(nums, left, right)

		// 递归排序左右两部分
		quickSortHelper(nums, left, p-1)
		quickSortHelper(nums, p+1, right)
	}
}

// 三数取中法选择基准
func medianOfThree(nums []int, left, right int) int {
	mid := left + (right-left)/2

	// 确保 nums[left] <= nums[mid] <= nums[right]
	if nums[left] > nums[mid] {
		nums[left], nums[mid] = nums[mid], nums[left]
	}
	if nums[mid] > nums[right] {
		nums[mid], nums[right] = nums[right], nums[mid]
	}
	if nums[left] > nums[mid] {
		nums[left], nums[mid] = nums[mid], nums[left]
	}

	return mid
}

// 双指针分区法（更简洁的实现）
func partition(nums []int, left, right int) int {
	pivot := nums[left]
	i, j := left+1, right

	for i <= j {
		// 从左找第一个大于pivot的
		for i <= right && nums[i] <= pivot {
			i++
		}
		// 从右找第一个小于等于pivot的
		for j > left && nums[j] > pivot {
			j--
		}
		// 交换
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	// 将pivot放到正确位置
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

// 小数组使用插入排序
func insertionSort(nums []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		key := nums[i]
		j := i - 1
		for j >= left && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

// 另一种分区实现：单指针法（类似你的partition但更简洁）
func partitionSinglePointer(nums []int, left, right int) int {
	pivot := nums[right] // 选最后一个作为基准
	i := left - 1        // i指向小于pivot部分的最后一个元素

	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

// 随机快排版本
func quickSortRandom(nums []int) {
	quickSortRandomHelper(nums, 0, len(nums)-1)
}

func quickSortRandomHelper(nums []int, left, right int) {
	if left < right {
		// 随机选择基准
		randomIndex := left + rand.Intn(right-left+1)
		nums[left], nums[randomIndex] = nums[randomIndex], nums[left]

		p := partition(nums, left, right)
		quickSortRandomHelper(nums, left, p-1)
		quickSortRandomHelper(nums, p+1, right)
	}
}
