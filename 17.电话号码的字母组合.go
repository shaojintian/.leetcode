/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
// @lc code=start
func letterCombinations(digits string) []string {
	res := make([]string, 0)
	l := len(digits)
	if l == 0 {
		return res
	}
	m := map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
	var backtrack func(combination, subDigits string)
	backtrack = func(combination, subDigits string) {
		if len(subDigits) == 0 {
			res = append(res, combination)
			return
		}
		for _, letter := range m[string(subDigits[0])] {
			combination_new := combination + string(letter)
			backtrack(combination_new, subDigits[1:])
		}
	}
	backtrack("", digits)
	return res
}

// @lc code=end

