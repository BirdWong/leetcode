package main

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}
	t := x
	var n int = 0
	var j int = 1
	for t != 0 {

		n = n*10 + t%10
		j = j * 10
		t = t / 10
		if t/j == 0 {
			break
		}
		if t/j < 10 {
			t = t / 10
			break
		}

	}
	return n == t

}

// @lc code=end
