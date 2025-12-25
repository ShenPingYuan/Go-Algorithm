package main

func main() {
	//[[1,4,5],[1,3,4],[2,6]]

	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	head3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	}
	var lists []*ListNode = []*ListNode{head1, head2, head3}
	mergeKLists(lists)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var listNodes = []*ListNode{}

	for i := 0; i < len(lists); i++ {
		list := lists[i]
		for list != nil {
			listNodes = append(listNodes, list)
			list = list.Next
		}
	}
	mergeSort(listNodes)

	dummy := &ListNode{}
	p := dummy
	for _, n := range listNodes {
		n.Next = nil
		p.Next = n
		p = p.Next
	}
	return dummy.Next
}

func mergeSort(nodes []*ListNode) {
	len := len(nodes)

	if len <= 1 {
		return
	}
	sort(nodes, 0, len-1)
}

func sort(nodes []*ListNode, start, end int) {
	if start >= end {
		return
	}
	mid := (end + start) / 2
	sort(nodes, start, mid)
	sort(nodes, mid+1, end)
	merge(nodes, start, mid, end)
}

func merge(nodes []*ListNode, start, mid, end int) {
	var temp []*ListNode

	p1, p2 := start, mid+1
	for p1 <= mid && p2 <= end {
		if nodes[p1].Val < nodes[p2].Val {
			temp = append(temp, nodes[p1])
			p1++
		} else {
			temp = append(temp, nodes[p2])
			p2++
		}
	}
	if p1 <= mid {
		temp = append(temp, nodes[p1:mid+1]...)
	}
	if p2 <= end {
		temp = append(temp, nodes[p2:end+1]...)
	}

	for i := 0; i < len(temp); i++ {
		nodes[start+i] = temp[i]
	}
}
