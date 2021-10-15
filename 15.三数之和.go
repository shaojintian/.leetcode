/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
 // map c = -(a+b)  map[c]index of c
 import sort
// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	if len(nums) <= 2 {
		return res
	}
	m := make(map[int]int, 0)
	for i, v := range nums {
		m[v] = i
	}
	
	for i := 0; i < len(nums); i++ {
		if i >=1 && nums[i] == nums[i-1]{
			continue
		}
		for k := i + 1; k < len(nums); k++ {
			if k > i+1 && nums[k] ==nums[k-1]{continue}
			c := -(nums[i] + nums[k])
			if cIndex, ok := m[c]; ok {
				if cIndex > k {
					tuple := []int{nums[i], nums[k], c}
					res = append(res, tuple)
				}
			}
		}
	}

	return res
}

// @lc code=end

