package main

/*
*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
3. 无重复字符的最长子串
中等
相关标签
premium lock icon
相关企业
提示
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/
func main() {

}

// 滑动窗口+双指针
func lengthOfLongestSubstring(s string) int {
	strLst := []rune(s)

	posMap := make(map[rune]int)
	maxLen := 0
	start := 0

	for i, r := range strLst {

		if pos, exist := posMap[r]; exist {
			start = pos + 1
		}

		posMap[r] = i

		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}

	}

	return maxLen

}
