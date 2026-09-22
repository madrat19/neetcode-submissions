func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Count, s2Count := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]]++
	}
	matches := 0
	for i := 0; i < len(s1); i++ {
		s2Count[s2[i]]++
		if num1, ok := s1Count[s2[i]]; ok && num1 >= s2Count[s2[i]] {
			matches++
		}
	}
	if matches == len(s1) {
		return true
	}
	for i, j := 0, len(s1); j < len(s2); i, j = i+1, j+1 {
		s2Count[s2[i]]--
		if num1, ok := s1Count[s2[i]]; ok && num1 > s2Count[s2[i]] {
			matches--
		}
		s2Count[s2[j]]++
		if num1, ok := s1Count[s2[j]]; ok && num1 >= s2Count[s2[j]] {
			matches++
		}
		if matches == len(s1) {
			return true
		}
	}
	return false
}
