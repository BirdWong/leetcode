package main

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	ansi := 0
	ansk := 0
	l := len(s)
	if l == 1 {
		return s
	}

	for i := 0; i < l-1; i++ {
		if l-i < ansk-ansi+1 {
			break
		}

		c := s[i]
		for k := l - 1; k > i; k-- {
			if k-i+1 < ansk-ansi+1 {
				break
			}
			if s[k] == c {
				ok := true
				r := k
				f := i
				for f < r {
					if s[f] == s[r] {
						f++
						r--
					} else {
						ok = false
						break
					}
				}

				if ok {
					ansk = k
					ansi = i
				}

			}

		}
	}
	return s[ansi : ansk+1]
}

// @lc code=end
