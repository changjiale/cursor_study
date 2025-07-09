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
