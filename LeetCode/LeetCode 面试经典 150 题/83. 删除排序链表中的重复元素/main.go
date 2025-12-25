package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	for p != nil && p.Next != nil {
		if p.Next.Val == p.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil {
		if fast.Val != slow.Val {
			// nums[slow] = nums[fast];
			slow.Next = fast
			// slow++;
			slow = slow.Next
		}
		// fast++
		fast = fast.Next
	}
	// 断开与后面重复元素的连接
	slow.Next = nil
	return head
}
