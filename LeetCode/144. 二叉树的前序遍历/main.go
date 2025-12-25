package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{root.Val}

	leftRes := preorderTraversal(root.Left)
	rightRes := preorderTraversal(root.Right)

	if leftRes != nil {
		res = append(res, leftRes...)
	}
	if rightRes != nil {
		res = append(res, rightRes...)
	}

	return res
}

func preorderTraversal1(root *TreeNode) []int {
	res := []int{}

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)
	return res
}
