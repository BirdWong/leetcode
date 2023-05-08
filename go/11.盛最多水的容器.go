/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
// min(height[i], height[j]) * (j-i)
package main

func maxArea(height []int) int {
	var max int = -1

	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j > i; j-- {
			if height[i]*(j-i) < max {
				break
			}
			if j < len(height)-1 && height[j] < height[j+1] {
				continue
			}

			m := Min(height[i], height[j]) * (j - i)

			if m > max {
				max = m
			}
		}
	}
	return max
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// @lc code=end
