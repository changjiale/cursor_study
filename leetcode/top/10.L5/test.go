package main

/*
*
中心扩散 以每个字符串往外部扩散 区分基数偶数
*/
func isP(s string) string {

	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := around(s, i, i)
		len2 := around(s, i, i+1)
		maxLen := max(len1, len2)
		if maxLen > end-start+1 {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

// 扩散长度
func around(s string, left int, right int) int {
	for left >= 0 && right <= len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1

}
