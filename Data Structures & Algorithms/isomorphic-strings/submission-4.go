func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false 
	}
	helper := func(s, t string) bool {
		mp := map[byte]byte{}
		for i := 0; i < len(s); i++ {
			if _, ok := mp[s[i]]; !ok {
				mp[s[i]] = t[i]
			}
			if mp[s[i]] != t[i] {
				return false 
			}
		}
		return true
	}
	return helper(s, t) && helper(t, s)
}
