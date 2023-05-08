package main

/*
 * @lc app=leetcode.cn id=1572 lang=golang
 *
 * [1572] 矩阵对角线元素的和
 */

// @lc code=start
func diagonalSum(mat [][]int) int {
	var sum int = 0
	length := len(mat)
	if length == 1 {
		return mat[0][0]
	}
	for i := 0; i < len(mat); i++ {
		sum += mat[i][i] + mat[i][length-1-i]
	}

	if length%2 == 1 {
		sum -= mat[length/2][length/2]
	}
	return sum
}

// @lc code=end
