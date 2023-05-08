/*
 * @lc app=leetcode.cn id=2370 lang=golang
 *
 * [2370] 最长理想子序列
 */
package main
// @lc code=start
func longestIdealString(s string, k int) int {
	var f []int = make([]int, 26)
    max := func(a, b int) int {
        if a >= b {
            return a
        }
        return b
    }
    min := func(a, b int) int {
        if a <= b {
            return a
        }
        return b
    }
    var m int = -1
	//var cm string
	for _, c := range s {
		v := int(c - 'a')
		for _, i := range f[max(v-k, 0):min(v+k+1, 26)] {
			f[v] = max(f[v], i)
		}
		f[v]++
		m = max(m, f[v])
		//cm = string(c)
		//println(cm, m)
	}
	//print(m, cm)
	return m

}
// @lc code=end



