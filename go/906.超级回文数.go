package main

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=906 lang=golang
 *
 * [906] 超级回文数
 */

// @lc code=start
func superpalindromesInRange(left string, right string) int {
	leftVal, _ := strconv.ParseInt(left, 10, 64)

	rightVal, _ := strconv.ParseFloat(right, 64)

	r := math.Floor(math.Sqrt(rightVal))

	count := 0

	for i := 1; i <= int(r); i++ {
		//println(i)
		strT := strconv.FormatInt(int64(i), 10)

		strB := []byte(strT)

		for k := len(strT) - 1; k >= 0; k-- {
			strB = append(strB, strT[k])
		}

		num1, _ := strconv.ParseInt(string(strB), 10, 64)

		strB = []byte(strT)
		for k := len(strT) - 2; k >= 0; k-- {
			strB = append(strB, strT[k])
		}

		num2, _ := strconv.ParseInt(string(strB), 10, 64)

		if num1 > int64(r) && num2 > int64(r) {
			break
		}

		num1 *= num1

		num2 *= num2

		if reverse(num1) == num1 && num1 <= int64(rightVal) && num1 >= leftVal {
			count++
		}

		if reverse(num2) == num2 && num2 <= int64(rightVal) && num2 >= leftVal {
			count++
		}

	}
	return count

}

func reverse(n int64) int64 {
	var ans int64 = 0
	for n > 0 {
		ans = 10*ans + n%10
		n /= 10
	}
	return ans
}

// @lc code=end
