/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 || l == 1 {
		return l
	}
	leftPtr := 0
	maxRes := 1
	for leftPtr < l-1 {
		res := 1
		m := make(map[byte]struct{}, 0)
		m[s[leftPtr]] = struct{}{}
		rightPtr := leftPtr + 1
		for rightPtr < l {
			_, ok := m[s[rightPtr]]
			if !ok {
				m[s[rightPtr]] = struct{}{}
				res++
			} else {
				break
			}
			rightPtr++
		}
		if res > maxRes {
			maxRes = res
		}
		leftPtr++
		if maxRes >= (l - 1 - leftPtr + 1) {
			break
		}
	}
	return maxRes
}

// @lc code=end

