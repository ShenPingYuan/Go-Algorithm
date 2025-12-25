package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 7, 3, 6, 2, 6, 5, 7, 3, 2, 1, 45, 34, 52, 346, 54, 63, 45, 4, 6, 357, 34, 5234, 525}
	heapSort(nums)
	fmt.Println(nums)

}

func heapSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		Swin(nums, i)
	}

	for i := len(nums) - 1; i >= 0; i-- {
		// 将最大的头换到最后
		Swap(nums, 0, i)
		// 再将新头下沉
		Sink(nums, 0, i-1)
	}
}

// func Pop() int {
// 	if h.Size < 0 {
// 		panic("error")
// 	}
// 	var head int = h.Heap[0]

// 	h.Heap[0] = h.Heap[h.Size-1]
// 	h.Size--
// 	h.Sink(0)
// 	return head
// }

func Parent(node int) int {
	if node < 1 {
		return -1
	}
	return (node - 1) / 2
}

func Left(node int) int {
	return 2*node + 1
}

func Right(node int) int {
	return 2*node + 2
}

func Swap(nums []int, x, y int) {
	nums[x], nums[y] = nums[y], nums[x]
}

func Swin(nums []int, node int) {
	for Parent(node) >= 0 && nums[node] > nums[Parent(node)] {
		Swap(nums, Parent(node), node)
		node = Parent(node)
	}
}

func Sink(nums []int, node int, lastIndex int) {
	for (Left(node) <= lastIndex && nums[node] < nums[Left(node)]) || (Right(node) <= lastIndex && nums[node] < nums[Right(node)]) {
		// 左右子节点都存在
		if Right(node) <= lastIndex {
			if nums[Left(node)] < nums[Right(node)] {
				Swap(nums, node, Right(node))
				node = Right(node)
			} else {
				Swap(nums, node, Left(node))
				node = Left(node)
			}
		} else { // 只存在左子节点
			Swap(nums, node, Left(node))
			node = Left(node)
		}
	}
}
