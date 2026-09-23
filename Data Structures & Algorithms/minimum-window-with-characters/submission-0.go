func minWindow(s string, t string) string {
	window, counter := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		counter[t[i]]++
	}
	ans := ""
	matches := 0
	i := 0
	for j := 0; j < len(s); j++ {
		right := s[j]
		window[right]++
		if num := counter[right]; num >= window[right] {
			matches++
		}
		if matches == len(t) {
			for matches == len(t) {
				left := s[i]
				window[left]--
				if num := counter[left]; num > window[left] {
					matches--
					if len(ans) > len(s[i:j+1]) || ans == "" {
						ans = s[i : j+1]
					}
				}
				i++
			}
		}

	}
	return ans
}
