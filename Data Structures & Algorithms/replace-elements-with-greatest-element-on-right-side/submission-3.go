func replaceElements(arr []int) []int {
	mx := -1
	ans := make([]int, len(arr))
	for i := len(arr)-1; i > -1; i-- {
		ans[i] = mx
		mx = max(mx, arr[i])
	}
	return ans
}
