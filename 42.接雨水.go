/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	var left, leftMax int = 0, height[0]
	var right, rightMax int = len(height) - 1, height[len(height) - 1]
	var sum int = 0
	for left < right {
		if leftMax < rightMax {
			left++
			if leftMax > height[left]{
				sum += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
		} else {
			right-- 
			if rightMax > height[right] {
				sum += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
		}
	}
	return sum
}
// @lc code=end

