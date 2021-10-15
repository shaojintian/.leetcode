/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	//init 0,n situation
	for col := 1; col <= len(p); col++ {
		if p[col-1] == '*' {
			dp[0][col] = dp[0][col-2]
		}
	}

	for row := 1; row <= len(s); row++ {
		for col := 1; col <= len(p); col++ {
			if s[row-1] == p[col-1] || p[col-1] == '.' {
				dp[row][col] = dp[row-1][col-1]
			} else {
				if p[col-1] == '*' {
					if dp[row][col-2] {
						dp[row][col] = true
					} else {
						if p[col-1-1] == s[row-1] || p[col-1-1] == '.' {
							dp[row][col] = dp[row-1][col]
						}
					}
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

// @lc code=end

