func getConcatenation(nums []int) []int {
    ans := []int{}
	for _, e := range nums {
		ans = append(ans, e)
	}
	for _, e := range nums {
		ans = append(ans, e)
	}
	return ans
}
