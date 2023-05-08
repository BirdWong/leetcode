pac
/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	tNode := &ListNode{Next: head}
	pre := tNode
	var tHead *ListNode

	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	tHead = pre.Next
	for i := left; i < right; i++ {
		next := tHead.Next
		tHead.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return tNode.Next
}

// @lc code=end

