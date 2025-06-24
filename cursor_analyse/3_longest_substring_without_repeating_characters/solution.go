package main

import (
	"fmt"
)

// 解法一：滑动窗口 + 哈希表（支持Unicode）
// 时间复杂度：O(n)，空间复杂度：O(min(m, n))，其中m是字符集大小
func lengthOfLongestSubstring1(s string) int {
	runes := []rune(s)
	charMap := make(map[rune]int)
	maxLen := 0
	start := 0

	for i, char := range runes {
		// 如果当前字符已经在窗口中出现过，更新start位置
		if lastPos, exists := charMap[char]; exists && lastPos >= start {
			start = lastPos + 1
		}
		// 更新当前字符的位置
		charMap[char] = i
		// 更新最大长度
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}

// 解法二：滑动窗口 + 数组（仅适用于ASCII字符）
// 时间复杂度：O(n)，空间复杂度：O(1)
func lengthOfLongestSubstring2(s string) int {
	// 使用数组记录每个字符最后出现的位置
	lastPos := [128]int{}
	for i := range lastPos {
		lastPos[i] = -1
	}

	maxLen := 0
	start := 0

	for i := 0; i < len(s); i++ {
		// 只处理ASCII字符
		if s[i] < 128 {
			if lastPos[s[i]] >= start {
				start = lastPos[s[i]] + 1
			}
			lastPos[s[i]] = i
			if i-start+1 > maxLen {
				maxLen = i - start + 1
			}
		}
	}
	return maxLen
}

// 解法三：滑动窗口 + 双指针（支持Unicode）
// 时间复杂度：O(n)，空间复杂度：O(min(m, n))
func lengthOfLongestSubstring3(s string) int {
	runes := []rune(s)
	charMap := make(map[rune]int)
	maxLen := 0
	start := 0

	for i, char := range runes {
		// 如果当前字符已经在窗口中出现过，更新start位置
		if lastPos, exists := charMap[char]; exists && lastPos >= start {
			start = lastPos + 1
		}
		// 更新当前字符的位置
		charMap[char] = i
		// 更新最大长度
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}

func lengthOfLongestSubstringSimple(s string) int {
	m, res, left := map[rune]int{}, 0, 0
	for i, c := range []rune(s) {
		if m[c] > left {
			left = m[c]
		}
		if i-left+1 > res {
			res = i - left + 1
		}
		m[c] = i + 1
	}
	return res
}

func main() {
	cases := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
		"a",
		"你好世界",
	}
	fmt.Println("lengthOfLongestSubstring1:")
	for _, s := range cases {
		fmt.Printf("输入: %q, 输出: %d\n", s, lengthOfLongestSubstring1(s))
	}
	fmt.Println("\nlengthOfLongestSubstring2 (仅ASCII):")
	for _, s := range cases {
		fmt.Printf("输入: %q, 输出: %d\n", s, lengthOfLongestSubstring2(s))
	}
	fmt.Println("\nlengthOfLongestSubstring3:")
	for _, s := range cases {
		fmt.Printf("输入: %q, 输出: %d\n", s, lengthOfLongestSubstring3(s))
	}
	fmt.Println("\nlengthOfLongestSubstringSimple:")
	for _, s := range cases {
		fmt.Printf("输入: %q, 输出: %d\n", s, lengthOfLongestSubstringSimple(s))
	}
}
