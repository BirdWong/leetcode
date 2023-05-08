/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var nodeList []*TreeNode
	var lastNode *TreeNode
	for root != nil || len(nodeList) != 0 {
		for root != nil {
			nodeList = append(nodeList, root)
			root = root.Left
		}

		root = nodeList[len(nodeList) - 1]
		if root.Right == nil || root.Right == lastNode {
			ans = append(ans, root.Val)
			nodeList = nodeList[:len(nodeList) - 1]
			lastNode = root
			root = nil
		} else {
			//nodeList = append(nodeList, root)
			root = root.Right
		}
	}
	return ans
}
// @lc code=end

