package Leetcode

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val int
     	Left *TreeNode
     	Right *TreeNode
}

// recursive solution
// will write an iterative solution
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	} else {
		return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
	}
}