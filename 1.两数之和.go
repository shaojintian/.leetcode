/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	m := make(map[int]int, 0)
	for i, num := range nums {
		if index_0, ok := m[target-num]; ok {
			res[0] = index_0
			res[1] = i
			return res
		} else {
			m[num] = i
		}
	}
	return res
}

// @lc code=end

