package main

/*
 * @lc app=leetcode.cn id=1160 lang=golang
 *
 * [1160] 拼写单词
 */

// @lc code=start
func countCharacters(words []string, chars string) int {

	charsBytes := []byte(chars)
	mat := make(map[byte]int)
	for _, charByte := range charsBytes {
		if val, ok := mat[charByte]; ok {
			mat[charByte] = val + 1
		} else {
			mat[charByte] = 1
		}
	}

	length := 0
	for _, word := range words {
		wordBytes := []byte(word)
		tmpMat := copy(mat)
		success := true
		for _, wordChar := range wordBytes {
			if val, ok := tmpMat[wordChar]; ok && val > 0 {
				tmpMat[wordChar] = val - 1
			} else {
				success = false
				break
			}
		}

		if success {
			length += len(word)
		}
	}
	return length
}

func copy(mat map[byte]int) map[byte]int {
	c := make(map[byte]int)
	for k, v := range mat {
		c[k] = v
	}
	return c

}

// @lc code=end
