package main

func main() {
	// [1,2,2]
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	newHead := deleteDuplicates1(head)
	// 打印新链表
	for newHead != nil {
		print(newHead.Val, " ")
		newHead = newHead.Next
	}
	// 输出：1
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := map[int]int{}
	p := head
	for p != nil {
		m[p.Val]++
		p = p.Next
	}

	dummy := &ListNode{}
	p1 := dummy
	p = head
	for p != nil {
		if m[p.Val] == 1 {
			p1.Next = p
			p1 = p1.Next
		}
		p = p.Next
		p1.Next = nil
	}
	return dummy.Next
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dUniq, dDup := &ListNode{}, &ListNode{Val: 101}
	p1, p2 := dUniq, dDup
	p := head
	for p != nil {
		if p.Next != nil && p.Val == p.Next.Val || p.Val == p2.Val {
			p2.Next = p
			p2 = p2.Next
		} else {
			p1.Next = p
			p1 = p1.Next
		}
		p = p.Next
		p1.Next = nil
		p2.Next = nil
	}
	return dUniq.Next
}

// 快慢双指针解法
func deleteDuplicates2(head *ListNode) *ListNode {
    dummy := &ListNode{Val: -1}
    p, q := dummy, head
    for q != nil {
        if q.Next != nil && q.Val == q.Next.Val {
            // 发现重复节点，跳过这些重复节点
            for q.Next != nil && q.Val == q.Next.Val {
                q = q.Next
            }
            q = q.Next
            // 此时 q 跳过了这一段重复元素
            if q == nil {
                p.Next = nil
            }
            // 不过下一段元素也可能重复，等下一轮 while 循环判断
        } else {
            // 不是重复节点，接到 dummy 后面
            p.Next = q
            p = p.Next
            q = q.Next
        }
    }
    return dummy.Next
}