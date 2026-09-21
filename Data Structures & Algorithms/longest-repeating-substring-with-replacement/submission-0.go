func characterReplacement(s string, k int) int {
	frequency := map[byte]int{}
	i, j := 0, 0
	mostF := 1
	ans := 1
	for j < len(s) {
		new := s[j]
		frequency[new]++
		if frequency[new] > mostF {
			mostF = frequency[new]
		}
		if j-i+1-mostF <= k {
			ans = max(ans, j-i+1)
		} else {
			for j-i+1-mostF > k {
				old := s[i]
				frequency[old]--
				i++
			}
			ans = max(ans, j-i+1)
		}
		j++
	}
	return ans
}
