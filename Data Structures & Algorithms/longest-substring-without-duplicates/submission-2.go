func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	ans := 0
	i, j := 0, 0
	chars := map[byte]int{}
	chars[s[0]] = 1
	for j < len(s)-1 {
		j++
		chars[s[j]]++
		if chars[s[j]] == 1 {
			ans = max(ans, j-i+1)
		} else {
			for chars[s[j]] != 1 {
				chars[s[i]]--
				i++
			}
			ans = max(ans, j-i+1)
		}
	}
	return ans
}