package main

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	var l int32 = 0
	var r int32 = int32(len(s)) - 1
	var lc byte
	var rc byte
	// 相遇则回文
	for l < r {
		lc = s[l]
		rc = s[r]
		// 非合规字符跳过
		for (lc < 97 || lc > 122) && (lc < 65 || lc > 90) && (lc < 48 || lc > 57) {
			if l >= r {
				return true
			}
			l++
			lc = s[l]
		}
		for (rc < 97 || rc > 122) && (rc < 65 || rc > 90) && (rc < 48 || rc > 57) {
			if l >= r {
				return true
			}
			r--
			rc = s[r]
		}
		// 转小写
		lc = lc | 0x20
		rc = rc | 0x20

		//println(l, lc, r, rc)
		// 不相等直接退出
		if rc != lc {
			return false
		}

		l++
		r--
	}

	return true
}

// @lc code=end
