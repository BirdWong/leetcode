package main

/*
 * @lc app=leetcode.cn id=1013 lang=golang
 *
 * [1013] 将数组分成和相等的三个部分
 */

// @lc code=start
func canThreePartsEqualSum(arr []int) bool {
	var sum int64 = 0
	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])
	}
	if sum%3 != 0 {
		return false
	}

	air := sum / 3
	var tmp int64 = 0
	count := 0
	for i := 0; i < len(arr); i++ {

		tmp += int64(arr[i])
		if count == 2 {
			continue
		}
		if tmp == air {
			tmp = 0
			count++
			if len(arr)-1 == i {
				return false
			}
		}
	}
	//println(count, tmp, air)
	if count == 2 && air == tmp {
		return true
	}

	return false
}

// @lc code=end
