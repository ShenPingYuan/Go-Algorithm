package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nums := []*ListNode{}
	for head != nil {
		nums = append(nums, head)
		head = head.Next
	}
	dummy := &ListNode{}
	p := dummy
	for i := len(nums) - 1; i >= 0; i-- {
		p.Next = nums[i]
		p = p.Next
		p.Next = nil
	}
	return dummy.Next
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	for head != nil {
		n := head
		head = head.Next
		n.Next = dummy.Next
		dummy.Next = n
	}

	return dummy.Next
}
func reverseList3(head *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	var traverse func(*ListNode)
	traverse = func(head *ListNode) {
		if head == nil {
			return
		}
		traverse(head.Next)
		p.Next = head
		p = p.Next
		p.Next = nil
	}
	traverse(head)
	return dummy.Next
}

func reverseList4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2 := head, head.Next
	head.Next = nil
	for p2 != nil {
		pn := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = pn
	}
	return p1
}

func reverseList5(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList5(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
