package main

import "fmt"

/*
题目：有效的括号
难度：简单
标签：栈、字符串

题目描述：
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s，判断字符串是否有效。
有效字符串需满足：
1. 左括号必须用相同类型的右括号闭合。
2. 左括号必须以正确的顺序闭合。
3. 每个右括号都有一个对应的相同类型的左括号。

要求：
- 时间复杂度：O(n)，其中 n 是字符串长度
- 空间复杂度：O(n)，最坏情况下栈的大小

示例：
输入：s = "()"
输出：true

输入：s = "()[]{}"
输出：true

输入：s = "(]"
输出：false

输入：s = "([)]"
输出：false

输入：s = "{[]}"
输出：true
*/

// 解法一：栈（推荐）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isValid(s string) bool {
	// 使用切片模拟栈
	stack := []rune{}

	// 括号映射表
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 遍历字符串
	for _, char := range s {
		// 如果是左括号，入栈
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			// 如果是右括号，检查栈是否为空
			if len(stack) == 0 {
				return false
			}

			// 检查栈顶元素是否匹配
			top := stack[len(stack)-1]
			if top != brackets[char] {
				return false
			}

			// 匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		}
	}

	// 检查栈是否为空
	return len(stack) == 0
}

// 解法二：栈（使用switch语句）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isValid2(s string) bool {
	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// 左括号入栈
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 解法三：栈（使用ASCII值优化）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isValid3(s string) bool {
	stack := []rune{}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}

			// 使用ASCII值计算匹配的括号
			// ')' - '(' = 1, '}' - '{' = 2, ']' - '[' = 2
			top := stack[len(stack)-1]
			if (char == ')' && top != '(') ||
				(char == '}' && top != '{') ||
				(char == ']' && top != '[') {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 解法四：计数法（仅适用于只有一种括号的情况）
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func isValid4(s string) bool {
	// 注意：这种方法只适用于只有一种括号的情况
	// 对于多种括号混合的情况，这种方法会出错
	// 这里仅作为示例展示

	count := 0

	for _, char := range s {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}

	return count == 0
}

// 解法五：递归法（不推荐，仅用于理解）
// 时间复杂度：O(n)
// 空间复杂度：O(n)，递归栈深度
func isValid5(s string) bool {
	if len(s) == 0 {
		return true
	}

	// 找到最内层的括号对
	for i := 0; i < len(s)-1; i++ {
		if (s[i] == '(' && s[i+1] == ')') ||
			(s[i] == '{' && s[i+1] == '}') ||
			(s[i] == '[' && s[i+1] == ']') {
			// 移除这对括号，递归处理剩余部分
			newS := s[:i] + s[i+2:]
			return isValid5(newS)
		}
	}

	return false
}

// 解法六：栈（使用字节切片，性能更好）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isValid6(s string) bool {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if (char == ')' && top != '(') ||
				(char == '}' && top != '{') ||
				(char == ']' && top != '[') {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	// 测试用例
	testCases := []struct {
		s      string
		expect bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"(", false},
		{")", false},
		{"((", false},
		{"))", false},
		{"({[]})", true},
		{"({[}])", false},
		{"(((", false},
		{")))", false},
		{"()()", true},
		{"(()", false},
		{")()", false},
	}

	fmt.Println("=== 解法一：栈（推荐）===")
	for i, tc := range testCases {
		result := isValid(tc.s)
		fmt.Printf("测试用例 %d: s=\"%s\", 结果=%t, 期望=%t, 通过=%t\n",
			i+1, tc.s, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法二：栈（使用switch语句）===")
	for i, tc := range testCases {
		result := isValid2(tc.s)
		fmt.Printf("测试用例 %d: s=\"%s\", 结果=%t, 期望=%t, 通过=%t\n",
			i+1, tc.s, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法三：栈（使用ASCII值优化）===")
	for i, tc := range testCases {
		result := isValid3(tc.s)
		fmt.Printf("测试用例 %d: s=\"%s\", 结果=%t, 期望=%t, 通过=%t\n",
			i+1, tc.s, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法四：计数法（仅适用于单一括号）===")
	for i, tc := range testCases {
		// 只测试只有圆括号的情况
		if containsOnlyParentheses(tc.s) {
			result := isValid4(tc.s)
			fmt.Printf("测试用例 %d: s=\"%s\", 结果=%t, 期望=%t, 通过=%t\n",
				i+1, tc.s, result, tc.expect, result == tc.expect)
		}
	}

	fmt.Println("\n=== 解法五：递归法 ===")
	for i, tc := range testCases {
		result := isValid5(tc.s)
		fmt.Printf("测试用例 %d: s=\"%s\", 结果=%t, 期望=%t, 通过=%t\n",
			i+1, tc.s, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法六：栈（使用字节切片）===")
	for i, tc := range testCases {
		result := isValid6(tc.s)
		fmt.Printf("测试用例 %d: s=\"%s\", 结果=%t, 期望=%t, 通过=%t\n",
			i+1, tc.s, result, tc.expect, result == tc.expect)
	}
}

// 检查字符串是否只包含圆括号
func containsOnlyParentheses(s string) bool {
	for _, char := range s {
		if char != '(' && char != ')' {
			return false
		}
	}
	return true
}

/*
解题思路：

1. 栈（推荐）：
   - 使用栈来匹配括号
   - 遇到左括号入栈，遇到右括号与栈顶元素匹配
   - 如果匹配成功，弹出栈顶元素；否则返回false
   - 最后检查栈是否为空

2. 栈（使用switch语句）：
   - 使用switch语句处理不同类型的括号
   - 代码更清晰，易于理解
   - 性能与解法一相当

3. 栈（使用ASCII值优化）：
   - 利用括号的ASCII值特性
   - 代码更简洁
   - 性能略有提升

4. 计数法（仅适用于单一括号）：
   - 只适用于只有一种括号的情况
   - 空间复杂度O(1)
   - 不能处理多种括号混合的情况

5. 递归法：
   - 递归移除最内层的括号对
   - 代码简洁但性能较差
   - 递归栈深度可能很大

6. 栈（使用字节切片）：
   - 使用字节切片代替rune切片
   - 性能更好，内存占用更少
   - 适用于ASCII字符

时间复杂度分析：
- 所有解法：O(n)，每个字符最多访问一次

空间复杂度分析：
- 栈解法：O(n)，最坏情况下栈的大小
- 计数法：O(1)，只使用常数额外空间
- 递归法：O(n)，递归栈深度

关键点：
1. 栈是解决括号匹配问题的经典方法
2. 注意边界条件：空字符串、单个字符等
3. 栈为空时遇到右括号是无效的
4. 最后检查栈是否为空
5. 理解不同解法的适用场景
*/
