package main

func lengthOfStr(s string) int {
	strMap := make(map[rune]int)
	start := 0
	maxLen := 0
	strList := []rune(s)

	for i, r := range strList {

		strMap[r] = i

		if pos, exist := strMap[r]; exist {
			start = pos + 1
		}

		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}

	}

	return maxLen

}
