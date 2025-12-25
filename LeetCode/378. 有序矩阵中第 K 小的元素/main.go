package main

import "container/heap"

func main() {
	// heap.Init()
}

type MinHeap []int

func kthSmallest(matrix [][]int, k int) int {
	h := &MinHeap{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			heap.Push(h, matrix[i][j])
		}
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}

func kthSmallest2(matrix [][]int, k int) int {
	h := &MinEHeap{}
	heap.Init(h)

	n := len(matrix)

	for i := 0; i < n; i++ {
		heap.Push(h, Element{
			val: matrix[i][0],
			row: i,
			col: 0,
		})
	}

	for i := 0; i < k-1; i++ {
		val := heap.Pop(h).(Element)
		if val.col < n-1 {
			heap.Push(h, Element{
				val: matrix[val.row][val.col+1],
				row: val.row,
				col: val.col + 1,
			})
		}
	}
	return heap.Pop(h).(Element).val
}

func (heap *MinHeap) Len() int {
	return len(*heap)
}

func left(node int) int {
	l := node*2 + 1
	if l < 0 {
		return (1 << 31) - 1
	}
	return l
}

func right(node int) int {
	return node*2 + 2
}

func parent(node int) int {
	return (node - 1) / 2
}

func (heap *MinHeap) Swap(a, b int) {
	(*heap)[a], (*heap)[b] = (*heap)[b], (*heap)[a]
}

func (heap *MinHeap) Push(n any) {
	(*heap) = append((*heap), n.(int))
	heap.flow(len(*heap) - 1)
}

func (heap *MinHeap) Pop() any {
	// head := (*heap)[0]
	// (*heap)[0] = (*heap)[len(*heap)-1] 
	// *heap = (*heap)[:len(*heap)-1]
	// heap.sink(0)
	// return head

	tail := (*heap)[len(*heap)-1]
	(*heap) = (*heap)[:len(*heap)-1]
	return tail
}

func (heap *MinHeap) Less(a, b int) bool {
	return (*heap)[a] < (*heap)[b]
}

func (heap *MinHeap) flow(n int) {
	for heap.Less(n, parent(n)) {
		heap.Swap(n, parent(n))
		n = parent(n)
	}
}

func (heap *MinHeap) sink(n int) {
	for left(n) < len(*heap) {
		min := n
		if right(n) < len(*heap) && heap.Less(right(n), min) {
			min = right(n)
		}

		if heap.Less(left(n), min) {
			min = left(n)
		}

		if min == n {
			break
		}

		heap.Swap(n, min)
		n = min
	}
}

type Element struct {
	val int
	row int
	col int
}

type MinEHeap []Element

func (h MinEHeap) Len() int {
	return len(h)
}

func (h MinEHeap) Less(a, b int) bool {
	return h[a].val < h[b].val
}

func (h *MinEHeap) Push(e any) {
	*h = append(*h, e.(Element))
}

func (h *MinEHeap) Pop() any {
	tail := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return tail
}

func (h MinEHeap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}
