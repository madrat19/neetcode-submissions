func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	for _, str := range strs {
		placed := false
		for i, arr := range ans {
			if isAnagram(str, arr[0]) {
				ans[i] = append(arr, str)
				placed = true
				break
			}

		}
		if !placed {
			ans = append(ans, []string{str})
		}
	}
	return ans
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	counter := make(map[rune]int)
	for _, r := range a {
		counter[r] += 1
	}

	for _, r := range b {
		counter[r] -= 1
	}
	for _, n := range counter {
		if n != 0 {
			return false
		}
	}
	return true
}
