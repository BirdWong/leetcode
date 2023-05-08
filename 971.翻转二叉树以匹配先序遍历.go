/*
 * @lc app=leetcode.cn id=971 lang=golang
 *
 * [971] 翻转二叉树以匹配先序遍历
 */

package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ans []int 
var tIdx int = 0
func f(r *TreeNode, resever bool, voyage []int) bool {
	
	if r == nil {
		return true
	}
	println(r.Val, tIdx, len(ans), resever)
	if r.Val == voyage[tIdx] {
		tIdx++
		if !f(r.Left, false, voyage) || !f(r.Right, false, voyage) {
			if !resever {
				ans = append(ans, r.Val)
				t := r.Right
				r.Right = r.Left
				r.Left = t
				return f(r, true, voyage)
			} else {
				tIdx--
				return false
			}
		}
		return true

	} else {
		tIdx--
		return false
	}
}
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	// lt 运行时直接重复调用 会导致公共变量累增
	ans = nil 
	tIdx = 0
	if f(root, false, voyage) {
		return ans
	} else {
		return []int{-1}
	}
}

// @lc code=end
