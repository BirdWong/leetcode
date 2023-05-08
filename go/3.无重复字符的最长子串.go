package main

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	numMap := make(map[byte]int)
	ans := 0
	start := 0
	for i := 0; i < len(s); i++ {
		c := s[i]

		if v, ok := numMap[c]; ok {
			start = Max(v, start)

		}

		ans = Max(i-start+1, ans)
		//println(c, start, i, ans)
		numMap[c] = i + 1

	}
	return ans

}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// @lc code=end
