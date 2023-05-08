package main

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	ts := ""
	l := len(s)
	if l <= numRows || numRows == 1 {
		return s
	}
	for k := 0; k < numRows; k++ {
		rev := false

		distance := 0
		for i := k; i < l; {
			ts = ts + string(s[i])

			distance, rev = reserve(numRows, k, rev)
			i = i + (numRows - distance - 1) + (numRows - 2 - distance) + 1

		}
	}
	return ts
}

func reserve(numRows, k int, r bool) (int, bool) {
	if k == 0 || k == numRows-1 {
		return 0, true
	}
	if r {
		return numRows - 1 - k, false
	} else {
		return k, true
	}

}

// @lc code=end
