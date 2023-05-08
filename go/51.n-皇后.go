package main

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start

var maps [][]string    // 记录答案
var defaultByte []byte // 初始化棋盘内容

func solveNQueens(n int) [][]string {

	maps = [][]string{} // 初始化存储

	var board []string
	// 初始化棋盘
	for i := 0; i < n; i++ {
		board = append(board, createBoard(n, -1))
	}
	// 运行
	scan(board, 0, n)

	return maps
}

func scan(board []string, row, n int) {
	// 当行数与要求大小相同时  说明遍历到最尾 直接存储
	if row == n {
		maps = append(maps, copy(board))
		return
	}

	// 遍历这行的所有点， 确认是否符合要求
	for col := 0; col < n; col++ {
		// 检查是否合规
		if checkAttack(board, row, col) {
			// 合规后在这行的col下标摆放棋子
			board[row] = createBoard(n, col)
			// 检查下一行（递归）
			scan(board, row+1, n)
			// 重置这一行内容
			board[row] = createBoard(n, -1)
		}
	}
}

func checkAttack(board []string, row, col int) bool {

	// 检查纵向是否存在皇后
	for i := 0; i < row; i++ {
		if byte(board[i][col]) == 'Q' {
			return false
		}
	}

	// 检查45度角是否存在皇后
	i := row - 1
	j := col - 1

	for i >= 0 && j >= 0 {
		if byte(board[i][j]) == 'Q' {
			return false
		}
		i--
		j--
	}

	// 检查135度角是否存在皇后
	i = row - 1
	j = col + 1

	for i >= 0 && j < len(board[0]) {
		if byte(board[i][j]) == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}

// 棋盘复制下来存储到map， 棋盘变量是指标， 直接保存会相互引用
func copy(src []string) []string {
	var targe []string
	targe = append(targe, src...)
	return targe
}

// 创建这一行， go string不可直接改变
func createBoard(n, col int) string {

	// 如果不需要摆放皇后， 初始化一次即可
	if col == -1 {
		if len(defaultByte) != n {
			for i := 0; i < n; i++ {
				defaultByte = append(defaultByte, '.')
			}
		}
		return string(defaultByte)
	}
	// 将皇后摆放在col位置
	angleBytes := []byte{}
	for i := 0; i < n; i++ {
		if i == col {
			angleBytes = append(angleBytes, 'Q')
		} else {
			angleBytes = append(angleBytes, '.')
		}

	}
	return string(angleBytes)
}

// @lc code=end
