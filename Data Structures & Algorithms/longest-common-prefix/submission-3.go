func longestCommonPrefix(strs []string) string {
    prefix := strs[0]
	for _, s := range strs {
		l := min(len(prefix), len(s))
		prefix = prefix[:l]
		for i := 0; i < l; i++ {
			if prefix[i] != s[i] {
			prefix = prefix[:i]
			break
			}
		}
	}
	return prefix
}
