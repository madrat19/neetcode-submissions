func maxDifference(s string) int {
	frequency := map[rune]int{}
	maxOdd := 0
	minEven := 101
	for _, r := range s {
		frequency[r]++
	}
	for _, f := range frequency {
		if f % 2 == 0 {
			minEven = min(minEven, f)
		} else {
			maxOdd = max(maxOdd, f)
		}
	}
	return maxOdd - minEven
}
