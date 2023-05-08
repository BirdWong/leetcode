/*
 * @lc app=leetcode.cn id=866 lang=golang
 *
 * [866] 回文素数
 */
package main

// @lc code=start
import "math"

func primePalindrome(n int) int {
	return findAll(n)
}

func findAll(n int) int {
	x := int(math.Log10(float64(n)))
	i := x
	if x%2 == 1 {
		i = x - 1
	}

	// var nums []int
	var lastA []int // 上批奇数
	reverse := func(num int) int {
		a := 0
		for num != 0 {
			a = (a * 10) + (num % 10)
			num = num / 10
		}
		return a
	}

	chushai := func(num int) bool {
		if num > 13 && (num%2 == 0 || num%3 == 0 || num%5 == 0 || num%7 == 0 || num%11 == 0 || num%13 == 0) {
			return false
		}
		return isS(num)
	}
	for ; i < 10; i++ {
		if i%2 == 0 {
			jinzhi := int(math.Pow10(i / 2))
			for k := jinzhi / 10; k < jinzhi; k++ {
				num := k*jinzhi + reverse(k)
				lastA = append(lastA, num)
				if i >= x && chushai(num) && num >= n {
					return num
					// nums = append(nums, num)
				}
			}
		} else {
			if i == 1 {
				for k := 0; k < 10; k++ {
					if chushai(k) && k >= n {
						return k
						// nums = append(nums, k)
					}
				}
			} else {
				jinzhi := int(math.Pow10(int(i-1) / 2))
				for _, j := range lastA {
					for k := 0; k < 10; k++ {
						num := j/jinzhi*jinzhi*10 + k*jinzhi + j%jinzhi
						if chushai(num) && num >= n {
							return num
							// nums = append(nums, n)
						}
					}
				}
			}
			lastA = []int{}
		}
	}
	return -1
}
func isS(num int) bool {
	if num < 2 {
		return false
	}
	i := int(math.Sqrt(float64(num)))
	for k := i; k > 1; k-- {
		if num%k == 0 {
			return false
		}
	}
	return true
}

// @lc code=end
