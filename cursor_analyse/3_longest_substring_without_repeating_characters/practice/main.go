package main

import (
	"fmt"
)

/*
题目：无重复字符的最长子串
难度：中等
标签：哈希表、字符串、滑动窗口

题目描述：
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

要求：
1. 子串必须是连续的
2. 不能包含重复字符
3. 返回最长子串的长度

示例：
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

提示：
- 使用滑动窗口算法
- 可以用哈希表记录字符位置
- 注意处理边界情况
*/

// TODO: 在这里实现你的算法
func lengthOfLongestSubstring(s string) int {
	// 请实现你的代码
	return 0
}

func main() {
	// 测试用例
	testCases := []struct {
		input  string
		output int
	}{
		{"abcabcbb", 3}, // 普通情况
		{"bbbbb", 1},    // 重复字符
		{"pwwkew", 3},   // 包含重复字符
		{"", 0},         // 空字符串
		{"a", 1},        // 单个字符
		{"你好世界", 4},     // Unicode字符
		{"au", 2},       // 两个不同字符
		{"aab", 2},      // 开头重复
		{"dvdf", 3},     // 中间重复
		{"anviaj", 5},   // 复杂情况
	}

	for _, tc := range testCases {
		fmt.Printf("\n输入: %q\n", tc.input)
		result := lengthOfLongestSubstring(tc.input)
		fmt.Printf("输出: %d\n", result)
		fmt.Printf("期望: %d\n", tc.output)
		if result == tc.output {
			fmt.Println("✓ 通过")
		} else {
			fmt.Println("✗ 失败")
		}
	}
}

/*
预期输出：
输入: "abcabcbb"
输出: 3
期望: 3
✓ 通过

输入: "bbbbb"
输出: 1
期望: 1
✓ 通过

输入: "pwwkew"
输出: 3
期望: 3
✓ 通过

输入: ""
输出: 0
期望: 0
✓ 通过

输入: "a"
输出: 1
期望: 1
✓ 通过

输入: "你好世界"
输出: 4
期望: 4
✓ 通过

输入: "au"
输出: 2
期望: 2
✓ 通过

输入: "aab"
输出: 2
期望: 2
✓ 通过

输入: "dvdf"
输出: 3
期望: 3
✓ 通过

输入: "anviaj"
输出: 5
期望: 5
✓ 通过
*/
