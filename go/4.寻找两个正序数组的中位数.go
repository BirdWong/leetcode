package main

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var num1point int = 0
	var num2point int = 0
	var mid int = (len(nums1) + len(nums2) + 1) / 2
	evenNumber := ((len(nums1)+len(nums2))&1 == 0)
	var val float64 = 0
	nums1end := false
	nums2end := false
	if len(nums2) == 0 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		}
		if evenNumber {
			return float64(nums1[mid]+nums1[mid-1]) / 2
		} else {
			return float64(nums1[mid-1])
		}

	}
	if len(nums1) == 0 {
		if len(nums2) == 1 {
			return float64(nums2[0])
		}
		if evenNumber {
			return float64(nums2[mid]+nums2[mid-1]) / 2
		} else {
			return float64(nums2[mid-1])
		}
	}

	var lastVal int = 0

	//println("mid: ", mid, "evenNum: ", evenNumber)
	for !nums1end || !nums2end {
		additional := 0
		if nums1end {
			additional += 1
		}
		if nums2end {
			additional += 1
		}
		//println("point1: ", num1point, "point2: ", num2point, "additional: ", additional, "lastVal: ", lastVal, "val: ", val)
		if num1point+num2point+additional == mid {
			if !evenNumber {
				//println("11")
				return float64(lastVal)
			}
			val = float64(lastVal)
		}

		if num1point+num2point+additional == mid+1 {
			val = (val + float64(lastVal)) / 2
			// println("22")

			return val
		}
		//println(nums1end, nums2end)
		if (nums1[num1point] <= nums2[num2point] && !nums1end) || nums2end {
			lastVal = nums1[num1point]

			if num1point == len(nums1)-1 {
				nums1end = true
			} else {
				num1point++
			}
		} else {
			lastVal = nums2[num2point]

			if num2point == len(nums2)-1 {
				nums2end = true
			} else {
				num2point++
			}
		}
	}
	//println("33")

	return (val + float64(lastVal)) / 2
}

// @lc code=end
