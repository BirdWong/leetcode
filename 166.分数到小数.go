/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 */
package main

// @lc code=start
import (
	"fmt"
	"math"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	subFlag := ""
	if numerator*denominator < 0 {
		subFlag = "-"
	}

	numerator = int(math.Abs(float64(numerator)))
	denominator = int(math.Abs(float64(denominator)))
	numStr := ""
	if numerator > denominator {
		numStr = fmt.Sprintf("%d", numerator/denominator)
		numerator = (numerator % denominator)
		if numerator != 0 {
			numStr = fmt.Sprintf("%s.", numStr)
		}
	} else {
		numStr = "0."
	}
	numerator *= 10

	var numerators []int = []int{numerator}

	for numerator != 0 {
		var value int = 0
		if numerator < denominator {
			numerator *= 10
		} else {
			value = numerator / denominator
			numerator = (numerator % denominator) * 10
		}
		numStr = fmt.Sprintf("%s%d", numStr, value)
		for i, v := range numerators {
			if v == numerator {
				fields := strings.Split(numStr, ".")
				numStr = fmt.Sprintf("%s.%s(%s)", fields[0], fields[1][:i], fields[1][i:])
				return subFlag + numStr
			}
		}
		numerators = append(numerators, numerator)

	}
	return subFlag + numStr
}

// @lc code=end
