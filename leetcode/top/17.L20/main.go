package main

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：

输入：s = "()"

输出：true

示例 2：

输入：s = "()[]{}"

输出：true

示例 3：

输入：s = "(]"

输出：false

示例 4：

输入：s = "([])"

输出：true
*/
func main() {

}


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
