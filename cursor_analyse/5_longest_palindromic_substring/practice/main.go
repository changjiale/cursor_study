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

示例：
输入: s = "babad"
输出: "bab"
解释: "aba" 同样是符合题意的答案。

输入: s = "cbbd"
输出: "bb"

输入: s = "a"
输出: "a"

提示：
- 可以使用动态规划
- 也可以使用中心扩展法
- 注意处理边界情况
*/

// TODO: 在这里实现你的算法
func longestPalindrome(s string) string {
	// 请实现你的代码
	return ""
}

func main() {
	// 测试用例
	testCases := []struct {
		input  string
		output string
	}{
		{"babad", "bab"},        // 普通情况
		{"cbbd", "bb"},          // 偶数长度回文
		{"a", "a"},              // 单个字符
		{"ac", "a"},             // 两个字符
		{"racecar", "racecar"},  // 完整回文
		{"abba", "abba"},        // 偶数长度完整回文
		{"abcba", "abcba"},      // 奇数长度完整回文
		{"", ""},                // 空字符串
		{"aaaa", "aaaa"},        // 全相同字符
		{"abacdfgdcaba", "aba"}, // 复杂情况
	}

	for _, tc := range testCases {
		fmt.Printf("\n输入: %q\n", tc.input)
		result := longestPalindrome(tc.input)
		fmt.Printf("输出: %q\n", result)
		fmt.Printf("期望: %q\n", tc.output)
		if result == tc.output {
			fmt.Println("✓ 通过")
		} else {
			fmt.Println("✗ 失败")
		}
	}
}

/*
预期输出：
输入: "babad"
输出: "bab"
期望: "bab"
✓ 通过

输入: "cbbd"
输出: "bb"
期望: "bb"
✓ 通过

输入: "a"
输出: "a"
期望: "a"
✓ 通过

输入: "ac"
输出: "a"
期望: "a"
✓ 通过

输入: "racecar"
输出: "racecar"
期望: "racecar"
✓ 通过

输入: "abba"
输出: "abba"
期望: "abba"
✓ 通过

输入: "abcba"
输出: "abcba"
期望: "abcba"
✓ 通过

输入: ""
输出: ""
期望: ""
✓ 通过

输入: "aaaa"
输出: "aaaa"
期望: "aaaa"
✓ 通过

输入: "abacdfgdcaba"
输出: "aba"
期望: "aba"
✓ 通过
*/
