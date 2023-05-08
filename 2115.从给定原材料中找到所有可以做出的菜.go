/*
 * @lc app=leetcode.cn id=2115 lang=golang
 *
 * [2115] 从给定原材料中找到所有可以做出的菜
 */
package main

// @lc code=start
// 暴力
// func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
// 	suppliesMap := make(map[string]bool, len(supplies))
// 	for _, supplie := range supplies {
// 		suppliesMap[supplie] = true
// 	}

// 	var deg []string
// 	var canDeg []int
// 	for i := range recipes {
// 		canDeg = append(canDeg, int(i))
// 	}
// 	lastSupplies := len(suppliesMap) + 1
// 	for len(suppliesMap) != lastSupplies {
// 		var tCanDeg []int
// 		lastSupplies = len(suppliesMap)
// 		for _, i := range canDeg {
			
// 			isOk := true
// 			for _, need := range ingredients[i] {
// 				if _, ok := suppliesMap[need]; !ok {
// 					isOk = false
// 					break
// 				}
// 			}
// 			if isOk {
// 				deg = append(deg, recipes[i])
// 				suppliesMap[recipes[i]] = true
// 			} else {
// 				tCanDeg = append(tCanDeg, i)
// 			}
// 		}
		
// 		canDeg = tCanDeg
// 	}
// 	return deg

// }


func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	deg := make(map[string]int, len(recipes))
	g := make(map[string][]string, len(recipes))

	for i := range recipes {
		for _, ingredient := range ingredients[i] {
			g[ingredient] = append(g[ingredient], recipes[i])
		}
		
		deg[recipes[i]] = len(ingredients[i])
	}


	var ans []string
	for ;len(supplies) > 0; {
		s := supplies[0]
		supplies = supplies[1:]
		for _, a := range g[s] {
			if deg[a]--; deg[a] == 0 {
				ans = append(ans, a)
				supplies = append(supplies, a)
			}
		}
	}
	return ans
}
// @lc code=end
