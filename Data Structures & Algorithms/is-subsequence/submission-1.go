func isSubsequence(s string, t string) bool {
	lens, lent := len(s), len(t)
	if lens == 0 {
		return true
	}
	i := 0
	for j := 0; j < lent; j++ {
		if s[i] == t[j] {
			i++
		}
		if i == lens {
			return true
		}
	}
	return false
}
