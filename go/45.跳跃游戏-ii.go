package main

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}

	freeTimes := nums[0] // 可以移动的步数
	idx := 0             // 当前索引
	times := 0           // 已经移动次数

	maxTimes := freeTimes // 可移动最大步数
	tIdx := idx           // 可移动最远距离时的索引
	for freeTimes > 0 {
		maxTimes = freeTimes
		tIdx = idx
		// 如果最大步数+最大步数时索引 大于等于数组最大索引， 直接跳入结束
		if maxTimes+tIdx >= length-1 {
			return times + 1
		}

		// 在最大步数范围内查找下一次的最大跳入范围， 同时注意越界
		for i := 1; i <= freeTimes && i+idx < length; i++ {

			// 如果数值 + 当前索引 >最大跳动次数+最大跳动次数时索引， 则更新最大跳动次数 和 最大索引id
			if nums[i+idx]+i+idx > maxTimes+tIdx {
				maxTimes = nums[i+idx]
				tIdx = i + idx
			}
			// 如果最大跳跃距离 + 最大跳跃索引 大于等于数组最大索引距离， 直接跳动两次结束
			if maxTimes+tIdx >= length-1 {
				return times + 2
			}

		}
		// 上述循环没有跳动到结尾， 跳动到最优解索引， 进入下次循环
		times += 1
		freeTimes = maxTimes
		idx = tIdx
	}
	return times

}

// func jump(nums []int) int {
// 	length := len(nums)
// 	if length == 1 {
// 		return 0
// 	}

// 	if nums[0] >= length {
// 		return 1
// 	}

// 	now := 0

// 	start := 0

// 	step := 0

// 	for i := 0; i < length; i++ {
// 		now = Max(now, i+nums[i])
// 		if now >= length-1 {
// 			return step + 1
// 		}

// 		if start == i {
// 			start = now
// 			step++
// 		}
// 	}
// 	return step
// }

// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end
