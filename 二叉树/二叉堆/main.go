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
	heap := NewMinHeap()
	for _, v := range nums {
		heap.Push(v)
	}

	for i := 0; i < len(nums); i++ {
		val := heap.Pop()
		nums[i] = val
	}
}

type MinHeap struct {
	Heap []int
	Size int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		Heap: []int{},
		Size: 0,
	}
}

func (h *MinHeap) Push(val int) {
	h.Heap = append(h.Heap, val)
	h.Size++
	h.Swin(h.Size - 1)
}

//jGZfXGm48yb0CBwUowqX0GI8k9KI4dy6

func (h *MinHeap) Pop() int {
	if h.Size < 0 {
		panic("error")
	}
	var head int = h.Heap[0]

	h.Heap[0] = h.Heap[h.Size-1]
	h.Size--
	h.Sink(0)
	return head
}

func (h *MinHeap) Parent(node int) int {
	if node < 1 {
		return -1
	}
	return (node - 1) / 2
}

func (h *MinHeap) Left(node int) int {
	return 2*node + 1
}

func (h *MinHeap) Right(node int) int {
	return 2*node + 2
}

func (h *MinHeap) Swap(x, y int) {
	h.Heap[x], h.Heap[y] = h.Heap[y], h.Heap[x]
}

func (h *MinHeap) Swin(node int) {
	for h.Parent(node) >= 0 && h.Heap[node] < h.Heap[h.Parent(node)] {
		h.Swap(h.Parent(node), node)
		node = h.Parent(node)
	}
}

func (h *MinHeap) Sink(node int) {
	for (h.Left(node) <= h.Size && h.Heap[node] > h.Heap[h.Left(node)]) || (h.Right(node) <= h.Size && h.Heap[node] > h.Heap[h.Right(node)]) {
		// 左右子节点都存在
		if h.Right(node) <= h.Size {
			if h.Heap[h.Left(node)] < h.Heap[h.Right(node)] {
				h.Swap(node, h.Left(node))
				node = h.Left(node)
			} else {
				h.Swap(node, h.Right(node))
				node = h.Right(node)
			}
		} else { // 只存在左子节点
			h.Swap(node, h.Left(node))
			node = h.Left(node)
		}
	}
}