package main

import (
	"fmt"
)

/*
题目：最长回文子串
难度：中等
标签：字符串、动态规划、中心扩展

题目描述：
给你一个字符串 s，找到 s 中最长的回文子串。

回文串：正着读和倒着读都一样的字符串。

要求：
1. 返回最长回文子串
2. 如果存在多个最长回文子串，返回任意一个
3. 字符串长度 >= 1
*/

// 解法一：暴力解法
// 时间复杂度：O(n³)
// 空间复杂度：O(1)
func longestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}

	maxLen := 1
	start := 0

	// 枚举所有可能的子串
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i+1 > maxLen && isPalindrome(s, i, j) {
				maxLen = j - i + 1
				start = i
			}
		}
	}

	return s[start : start+maxLen]
}

// 判断子串是否为回文
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 解法二：动态规划
// 时间复杂度：O(n²)
// 空间复杂度：O(n²)
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}

	n := len(s)
	// dp[i][j] 表示 s[i:j+1] 是否为回文串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// 所有单个字符都是回文串
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	maxLen := 1
	start := 0

	// 枚举子串长度
	for length := 2; length <= n; length++ {
		// 枚举左边界
		for i := 0; i <= n-length; i++ {
			j := i + length - 1 // 右边界

			if s[i] == s[j] {
				if length == 2 || dp[i+1][j-1] {
					dp[i][j] = true
					if length > maxLen {
						maxLen = length
						start = i
					}
				}
			}
		}
	}

	return s[start : start+maxLen]
}

// 解法三：中心扩展法
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
func longestPalindrome3(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0

	// 以每个字符为中心，向两边扩展
	for i := 0; i < len(s); i++ {
		// 奇数长度回文串
		len1 := expandAroundCenter(s, i, i)
		// 偶数长度回文串
		len2 := expandAroundCenter(s, i, i+1)

		maxLen := max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

// 从中心向两边扩展
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// 解法四：Manacher算法（马拉车算法）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func longestPalindrome4(s string) string {
	if len(s) < 2 {
		return s
	}

	// 预处理字符串，在字符间插入特殊字符
	t := "#"
	for _, char := range s {
		t += string(char) + "#"
	}

	n := len(t)
	p := make([]int, n) // p[i] 表示以i为中心的最长回文半径

	// 中心点和右边界
	center, right := 0, 0

	for i := 0; i < n; i++ {
		// 利用对称性
		if i < right {
			mirror := 2*center - i
			p[i] = min(right-i, p[mirror])
		}

		// 向两边扩展
		left := i - (p[i] + 1)
		right_bound := i + (p[i] + 1)
		for left >= 0 && right_bound < n && t[left] == t[right_bound] {
			p[i]++
			left--
			right_bound++
		}

		// 更新中心点和右边界
		if i+p[i] > right {
			center = i
			right = i + p[i]
		}
	}

	// 找到最长回文子串
	maxLen := 0
	centerIndex := 0
	for i, radius := range p {
		if radius > maxLen {
			maxLen = radius
			centerIndex = i
		}
	}

	// 计算原字符串中的起始位置
	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}

// 辅助函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 测试用例
	testCases := []string{
		"babad",        // 期望输出: "bab" 或 "aba"
		"cbbd",         // 期望输出: "bb"
		"a",            // 期望输出: "a"
		"ac",           // 期望输出: "a" 或 "c"
		"racecar",      // 期望输出: "racecar"
		"abba",         // 期望输出: "abba"
		"abcba",        // 期望输出: "abcba"
		"",             // 期望输出: ""
		"aaaa",         // 期望输出: "aaaa"
		"abacdfgdcaba", // 期望输出: "aba"
	}

	for _, s := range testCases {
		fmt.Printf("\n输入: %s\n", s)

		// 解法一：暴力解法
		result1 := longestPalindrome1(s)
		fmt.Printf("解法一（暴力解法）结果: %s\n", result1)

		// 解法二：动态规划
		result2 := longestPalindrome2(s)
		fmt.Printf("解法二（动态规划）结果: %s\n", result2)

		// 解法三：中心扩展法
		result3 := longestPalindrome3(s)
		fmt.Printf("解法三（中心扩展法）结果: %s\n", result3)

		// 解法四：Manacher算法
		result4 := longestPalindrome4(s)
		fmt.Printf("解法四（Manacher算法）结果: %s\n", result4)
	}
}

/*
预期输出：
输入: babad
解法一（暴力解法）结果: bab
解法二（动态规划）结果: bab
解法三（中心扩展法）结果: bab
解法四（Manacher算法）结果: bab

输入: cbbd
解法一（暴力解法）结果: bb
解法二（动态规划）结果: bb
解法三（中心扩展法）结果: bb
解法四（Manacher算法）结果: bb

输入: a
解法一（暴力解法）结果: a
解法二（动态规划）结果: a
解法三（中心扩展法）结果: a
解法四（Manacher算法）结果: a

输入: ac
解法一（暴力解法）结果: a
解法二（动态规划）结果: a
解法三（中心扩展法）结果: a
解法四（Manacher算法）结果: a

输入: racecar
解法一（暴力解法）结果: racecar
解法二（动态规划）结果: racecar
解法三（中心扩展法）结果: racecar
解法四（Manacher算法）结果: racecar

输入: abba
解法一（暴力解法）结果: abba
解法二（动态规划）结果: abba
解法三（中心扩展法）结果: abba
解法四（Manacher算法）结果: abba

输入: abcba
解法一（暴力解法）结果: abcba
解法二（动态规划）结果: abcba
解法三（中心扩展法）结果: abcba
解法四（Manacher算法）结果: abcba

输入:
解法一（暴力解法）结果:
解法二（动态规划）结果:
解法三（中心扩展法）结果:
解法四（Manacher算法）结果:

输入: aaaa
解法一（暴力解法）结果: aaaa
解法二（动态规划）结果: aaaa
解法三（中心扩展法）结果: aaaa
解法四（Manacher算法）结果: aaaa

输入: abacdfgdcaba
解法一（暴力解法）结果: aba
解法二（动态规划）结果: aba
解法三（中心扩展法）结果: aba
解法四（Manacher算法）结果: aba
*/

/*
算法分析：

1. 暴力解法：
   - 枚举所有可能的子串，判断是否为回文
   - 时间复杂度高，但思路简单直观

2. 动态规划：
   - 利用回文串的对称性质
   - 状态转移方程：dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
   - 空间复杂度较高

3. 中心扩展法：
   - 以每个字符为中心，向两边扩展
   - 需要考虑奇数长度和偶数长度两种情况
   - 空间复杂度最优

4. Manacher算法：
   - 线性时间复杂度的最优解法
   - 利用回文串的对称性质避免重复计算
   - 算法较为复杂，但性能最优
*/
