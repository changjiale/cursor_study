package main

// 括号匹配
func test(s string) bool {
	strList := []rune(s)

	stack := []rune{}
	strMap := map[rune]rune{
		')': '(',
		//....
	}
	for _, str := range strList {
		if str == '(' || str == '{' || str == '[' {
			stack = append(stack, str)
		} else {
			//获取栈顶元素  相等
			if strMap[str] == stack[len(stack)-1] {
				//出站
				stack = stack[:len(stack)-1]
			}

		}

	}
	return len(stack) == 0
}
