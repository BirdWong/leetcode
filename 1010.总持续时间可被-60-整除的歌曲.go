/*
 * @lc app=leetcode.cn id=1010 lang=golang
 *
 * [1010] 总持续时间可被 60 整除的歌曲
 */

// @lc code=start
func numPairsDivisibleBy60(time []int) int {
	modTime := make(map[int][]int)
	for i, value := range time {
		mod := value % 60 
		sclice := modTime[mod]
		sclice = append(sclice, i)
		modTime[mod] = sclice
	}

	var sum int = 0
	for i, value := range time {
		need := (60 - (value % 60)) % 60
		if sclice, has := modTime[need]; has {
			for k, index := range sclice {
				if index > i {
					sum += (len(sclice) - k)
					break
				}
			}
		}
	}
	return sum
}
// @lc code=end

