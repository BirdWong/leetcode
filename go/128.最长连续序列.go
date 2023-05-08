package main

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	nummap := make(map[int]int, len(nums))
	var mv int = 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if _, ok := nummap[num]; ok {
			continue
		}
		var lv int = 0
		var rv int = 0
		var cv int = 0

		if value, ok := nummap[num-1]; ok {
			lv = value
		}

		if value, ok := nummap[num+1]; ok {
			rv = value
		}

		cv = lv + rv + 1

		if cv > mv {
			mv = cv
		}

		nummap[num] = cv
		nummap[num-lv] = cv
		nummap[num+rv] = cv
	}

	return mv
}

// @lc code=end
