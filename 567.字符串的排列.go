/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */



// @lc code=start


func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	res := make(map[string]int)
	for _, c := range s1 {
		value, _ := res[string(c)]
		res[string(c)] = value+1
	}

	left := 0
	for right, c := range s2 {
		value, _ := res[string(c)]
		value -= 1
		res[string(c)] = value 
		for value < 0 {
			tValue, _ := res[string(s2[left])] 
			res[string(s2[left])] = tValue + 1
			left++
			value, _ = res[string(c)]
		}
	
		if right-left+1 == n {
			return true
		}
	}
	return false
	
}
// @lc code=end

