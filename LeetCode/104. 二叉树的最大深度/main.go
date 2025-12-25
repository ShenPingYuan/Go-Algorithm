package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	var maxDepth int
// 	var depth int = 0
// 	traverse(root, &depth, &maxDepth)
// 	return maxDepth
// }

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)

	return max(leftMaxDepth, rightMaxDepth) + 1
}

func traverse(root *TreeNode, depth *int, maxDepth *int) {
	if root == nil {
		return
	}
	*depth++
	// 遍历到叶子节点时记录最大深度
	if root.Left == nil && root.Right == nil {
		*maxDepth = max(*depth, *maxDepth)
	}
	traverse(root.Left, depth, maxDepth)
	traverse(root.Right, depth, maxDepth)
	*depth--
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
