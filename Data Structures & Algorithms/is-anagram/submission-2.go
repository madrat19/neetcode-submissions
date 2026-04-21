
func isAnagram(s string, t string) bool {
	hash := make(map[rune]int)
	for _, r := range s {
		hash[r] += 1
	}
	for _, r := range t {
		hash[r] -= 1
	}

	for _, n := range hash {
		if n != 0 {
			return false
		}
	}
	return true
}
