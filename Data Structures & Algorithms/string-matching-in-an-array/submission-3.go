func stringMatching(words []string) []string {
	mp := map[string]struct{}{}
	ans := []string{}
    for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j && strings.Contains(words[i], words[j]) {
				mp[words[j]] = struct{}{}
			}
		}
	}
	for k, _ := range mp {
		ans = append(ans, k)
	}
	return ans
}