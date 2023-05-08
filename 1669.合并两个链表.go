/*
 * @lc app=leetcode.cn id=1669 lang=golang
 *
 * [1669] 合并两个链表
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var root *ListNode = &ListNode{}
	var index *ListNode
	i := 0
	index = list1
	root = list1
	for list1 != nil {

		//println(i, index.Val)
		if i == a {
			if i == 0 {
				index = list2
				root = list2
			} else {
				index.Next = list2
			}
			for index.Next != nil {
				index = index.Next
			}

		}

		list1 = list1.Next
		i++
		if a <= i && b >= i {
			continue
		}

		index.Next = list1
		index = index.Next

	}
	return root
}

// @lc code=end
