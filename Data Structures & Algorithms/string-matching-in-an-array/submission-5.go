func stringMatching(words []string) []string {
	ans := []string{}
    for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j && strings.Contains(words[j], words[i]) {
				ans = append(ans, words[i])
				break
			}
		}
	}
	return ans
}