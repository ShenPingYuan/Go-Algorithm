package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	maxDepth(root, &res)
	return res
}

func maxDepth(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left, res)
	rightMax := maxDepth(root.Right, res)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	*res = max(*res, leftMax+rightMax)
	return max(leftMax, rightMax) + 1
}
