/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	//双指针
	p1 := 0
	p2 := len(height) - 1
	res := 0
	for p1 < p2 {
		tmpArea := min(height[p1], height[p2]) * (p2 - p1)
		if tmpArea > res {
			res = tmpArea
		}
		if height[p1] <= height[p2] {
			p1++
		} else {
			p2--
		}
	}
	return res
}
func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

// @lc code=end

