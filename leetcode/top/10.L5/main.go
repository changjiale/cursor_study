package main

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

//中心扩散
//暴力
//两个写法

// 暴力
func longestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}
	max := 1
	start := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i+1 > max && isPalindrome(s, i, j) {
				max = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+max]

}
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

//中心扩散
/**
1. 遍历字符串，以每个字符为中心，向两边扩散
2. 如果左右两边字符相同，则继续扩散
3. 如果左右两边字符不同，则停止扩散
4. 记录最长回文子串
5. 返回最长回文子串
*/
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		//判断回文串
		//基数回文
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := max(len1, len2)
		if maxLen > end-start+1 {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right <= len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}
