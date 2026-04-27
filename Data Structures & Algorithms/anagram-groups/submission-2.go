func groupAnagrams(strs []string) [][]string {
	res := map[[26]int][]string{}

	for _, str := range strs {
		counter := [26]int{}
		for _, c := range str {
			counter[c-'a'] += 1
		}
		res[counter] = append(res[counter], str)
	}

	ans := [][]string{}
	for _, group := range res {
		ans = append(ans, group)
	}

	return ans
}
