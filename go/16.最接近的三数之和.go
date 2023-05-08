package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	// fmt.Printf("%v\n", nums)
	length := len(nums)
	value := nums[0] + nums[1] + nums[length-1]
	for i := 0; i < length-2; i++ {

		l := i + 1
		k := length - 1

		for l < k {
			v := nums[i] + nums[l] + nums[k]
			// println(i, l, k)
			// println(v)
			if target-v == 0 {
				return v
			}

			if target > v {
				l++
			}

			if target < v {
				k--
			}

			if Max(target, value)-Min(target, value) > Max(target, v)-Min(target, v) {
				value = v
			}

		}
	}
	return value

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
