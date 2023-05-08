/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */

package main

// @lc code=start
func shipWithinDays(weights []int, days int) int {
	var left, right int = -1, 0
	for _, weight := range weights {
		if weight > left {
			left = weight
		}
		right += weight
	}
	var mid int = (left + right) / 2
	for left < right {

		tDays := int(1)
		tMid := mid
		for _, weight := range weights {
			if tMid < weight {
				tDays++
				tMid = mid
			}
			tMid -= weight
		}

		if tDays <= days {
			right = mid
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2

	}

	return mid

}

// @lc code=end
