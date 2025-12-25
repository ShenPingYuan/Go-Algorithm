package main

func main() {
	// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
	lists := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6}},
	}
	merged := mergeKLists(lists)
	// 打印merged
	for merged != nil {
		print(merged.Val, " ")
		merged = merged.Next
	}
	// 输出：merged = [1,1,2,3,4,4,5,6]
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{}
	p := dummy
	nums := make([]*ListNode, 0, len(lists))
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			push(&nums, lists[i])
		}
	}

	for len(nums) > 0 {
		node := pop(&nums)
		p.Next = node
		node = node.Next
		p = p.Next
		p.Next = nil
		if node != nil {
			push(&nums, node)
		}
	}
	return dummy.Next
}

func push(nums *[]*ListNode, node *ListNode) {
	*nums = append(*nums, node)

	current := len(*nums) - 1
	// 上浮
	for parent(current) != -1 && (*nums)[current].Val < (*nums)[parent(current)].Val {
		swap(*nums, current, parent(current))
		current = parent(current)
	}
}

func pop(nums *[]*ListNode) *ListNode {
	head := (*nums)[0]
	(*nums)[0] = (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]

	// 下沉
	currentIndex := 0
	for left(currentIndex) < len(*nums) {
		smaller := left(currentIndex)
		if right(currentIndex) < len(*nums) && (*nums)[right(currentIndex)].Val < (*nums)[left(currentIndex)].Val {
			smaller = right(currentIndex)
		}
		if (*nums)[currentIndex].Val <= (*nums)[smaller].Val {
			break
		}
		swap(*nums, currentIndex, smaller)
		currentIndex = smaller
	}

	return head
}

func swap(nums []*ListNode, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func left(index int) int {
	return index*2 + 1
}

func right(index int) int {
	return index*2 + 2
}

func parent(index int) int {
	if index == 0 {
		return -1
	}
	return (index - 1) / 2
}
