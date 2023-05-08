/*
 * @lc app=leetcode.cn id=816 lang=golang
 *
 * [816] 模糊坐标
 */
package main

// @lc code=start
import (
	"fmt"
)
func checkNum(str string, isDecimal int) bool {
	if len(str) < 2 {
		return true
	} 
	
	if isDecimal != -1 {
		if len(str) == 2 {
			return false
		}
		
		if str[len(str) - 1] == '0' {
			return false
		}
	} 
	
	if str[0] - '0' == 0 && str[1] - '0' == 0 {
		return false
	}

	if isDecimal != 0 && str[0] - '0' == 0 {
		return false
	}



	return true
}

func createNum(str string) []string {
	var nums []string 
	if checkNum(str, -1) {
		nums = append(nums, str)
	}
	for i := 1 ; i < len(str); i++ {
		s := fmt.Sprintf("%s.%s", str[:i] ,str[i:])
		if checkNum(s, i-1) {
			nums = append(nums, s)
		}
	}
	return nums
}

func ambiguousCoordinates(s string) []string {
	s = s[1:len(s)-1]
	var ans []string
	for i:=0; i < len(s) - 1 ; i++ {
		aNums := createNum(s[:i+1])
		if len(aNums) == 0 {
			continue
		}
		bNums := createNum(s[i+1:])
		for _, a := range aNums {
			for _, b := range bNums {
				ans = append(ans, fmt.Sprintf("(%s, %s)", a, b))
			}
		}
	}
	// if s[0] - '0' == 0 && s[1] - '0' != 0 {
	// 	for _, b := range createNum(s) {
	// 		ans = append(ans, fmt.Sprintf("(0, %s)", b))
	// 	}
	// }
	return ans
	// return fmt.Sprintf("[%s]",strings.Join(ans, ", "))
}
// @lc code=end

