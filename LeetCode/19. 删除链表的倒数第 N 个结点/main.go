package main

func main() {
	// 输入：head = [1,2,3,4,5], n = 2
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	n := 2
	newHead := removeNthFromEnd(head, n)
	// 打印新链表
	for newHead != nil {
		print(newHead.Val, " ")
		newHead = newHead.Next
	}
	// 输出：1 2 3 5
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p1, p2 := dummy, dummy
	for i := 0; i < n+1; i++ {
		p2 = p2.Next
	}
	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	pnn := p1.Next.Next
	p1.Next = pnn

	return dummy.Next
}
