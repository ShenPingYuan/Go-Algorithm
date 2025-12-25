package main

import (
	"fmt"
)

func main() {
	// Input array representing the binary tree
	rootArray := []interface{}{2, nil, 3, nil, 4, nil, 5, nil, 6}
	var tree = buildTree(rootArray, 0)
	traverseTree(tree)
	// fmt.Println(minDepth(tree))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Helper function to build the binary tree from a level-order array
func buildTree(arr []interface{}, index int) *TreeNode {
	// If index is out of bounds or the value is nil, return nil
	if index >= len(arr) || arr[index] == nil {
		return nil
	}

	// Create the current node
	node := &TreeNode{Val: arr[index].(int)}

	// Recursively build the left and right subtrees
	node.Left = buildTree(arr, 2*index+1)
	node.Right = buildTree(arr, 2*index+2)

	return node
}

func traverseTree(root *TreeNode) {
	if root == nil {
		return
	}
	var queue []*TreeNode = []*TreeNode{root}

	for len(queue) > 0 {
		var node = queue[0]
		queue = queue[1:]
		fmt.Println(node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

// 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth int = 1
	var queue []*TreeNode = []*TreeNode{root}

	for len(queue) > 0 {
		var currentLevelCount = len(queue)
		for i := 0; i < currentLevelCount; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
