package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nums1 := []int{}
	nums2 := []int{}
	dummy := &ListNode{}
	p := dummy
	for l1 != nil {
		nums1 = append(nums1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		nums2 = append(nums2, l2.Val)
		l2 = l2.Next
	}
	reverse(nums1)
	reverse(nums2)
	nums3 := []int{}
	carry := 0
	for i := 0; i < len(nums1) || i < len(nums2); i++ {
		v := carry
		carry = 0
		if i < len(nums1) {
			v += nums1[i]
		}
		if i < len(nums2) {
			v += nums2[i]
		}
		if v >= 10 {
			v -= 10
			carry = 1
		}
		nums3 = append(nums3, v)
	}
	if carry == 1 {
		nums3 = append(nums3, 1)
	}
	reverse(nums3)
	for i := 0; i < len(nums3); i++ {
		p.Next = &ListNode{Val: nums3[i]}
		p = p.Next
	}
	return dummy.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	nums1 := []int{}
	nums2 := []int{}
	dummy := &ListNode{}
	p := dummy.Next
	for l1 != nil {
		nums1 = append(nums1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		nums2 = append(nums2, l2.Val)
		l2 = l2.Next
	}
	carry := 0
	i1, i2 := len(nums1)-1, len(nums2)-1
	for i1 >= 0 || i2 >= 0 || carry == 1 {
		v := carry
		carry = 0
		if i1 >= 0 {
			v += nums1[i1]
			i1--
		}

		if i2 >= 0 {
			v += nums2[i2]
			i2--
		}

		if v >= 10 {
			v -= 10
			carry = 1
		}

		n := &ListNode{Val: v, Next: p}
		dummy.Next = n
		p = n
	}

	return dummy.Next
}

func reverse(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}
	start, end := 0, n-1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
