/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
// @lc code=start
func longestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	maxLen := 1
	maxString := string(s[0])
	dp := make([][]bool, len(s))
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for i := 0; i < l; i++ {
		for k := i + 1; k < l; k++ {
			ll := k - i + 1
			if ll == 2 && s[i] == s[k] {
				dp[i][k] = true
				if ll > maxLen {
					maxString = s[i : k+1]
					maxLen = ll
				}
			} else if k-i+1 == 2 && s[i] != s[k] {
				dp[i][k] = false
			}
			if k-i+1 >= 3 {
				if s[i] == s[k] && dp[i+1][k-1] == true {
					dp[i][k] = true
					if ll > maxLen {
						maxString = s[i : k+1]
						maxLen = ll
					}
				} else {
					dp[i][k] = false
				}
			}

		}
	}
	return maxString
}

// @lc code=end

