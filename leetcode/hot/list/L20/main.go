package L20

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func isValid1(s string) bool {
	if len(s)%2 != 0 { // s 长度必须是偶数
		return false
	}
	mp := map[byte]byte{')': '(', ']': '[', '}': '{'}
	st := []byte{}
	for i, _ := range s {
		c := byte(s[i])
		if mp[c] == 0 { // c 是左括号
			st = append(st, c) // 入栈
		} else { // c 是右括号
			if len(st) == 0 || st[len(st)-1] != mp[c] {
				return false // 没有左括号，或者左括号类型不对
			}
			st = st[:len(st)-1] // 出栈
		}
	}
	return len(st) == 0 // 所有左括号必须匹配完毕
}

func isValidtest(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	mp := map[byte]byte{
		')': '(',
	}
	st := []byte{}
	for i, _ := range s {
		t := s[i]
		if mp[t] == 0 { //左
			st = append(st, t)
		} else {
			if len(st) == 0 || st[len(st)-1] != mp[t] {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return true
}
