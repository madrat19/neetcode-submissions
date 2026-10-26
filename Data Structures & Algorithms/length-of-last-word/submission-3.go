func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for s[i] == ' ' {
		i--
	}
	if i == -1 {
		return 0
	}
	j := i
	for; j > -1 && s[j] != ' '; j-- {

	}
	return i - j
}
