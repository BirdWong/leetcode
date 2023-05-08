package main

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	var num int32 = 0
	var flag int32 = 100
	var ulimit int32 = 2147483647
	var dlimit int32 = -2147483648

	for i := 0; i < len(s); i++ {
		if s[i] < 48 || s[i] > 57 {
			if flag == 100 {
				if s[i] == ' ' {
					continue
				} else if s[i] == '-' {
					flag = -1
					continue
				} else if s[i] == '+' {
					flag = 1
					continue
				}
			}
			break

		}
		tn := s[i] - 48
		//println("tip: ", num, int(tn))

		if flag == -1 {
			if num == dlimit {
				continue
			}
			if num > dlimit/10 || num == 0 {
				num = num*10 - int32(tn)
			} else if num == dlimit/10 && int32(tn) < (dlimit%10)*-1 {
				num = num*10 - int32(tn)
			} else {
				num = dlimit
			}

		}
		if flag == 100 || flag == 1 {
			if flag == 100 {
				flag = 1
			}
			if num == ulimit {
				continue
			}
			if num < ulimit/10 {
				num = num*10 + int32(tn)
			} else if num == ulimit/10 && int32(tn) < (ulimit%10) {
				num = num*10 + int32(tn)
			} else {
				num = ulimit
			}

		}

	}
	return int(num)
}

// @lc code=end
