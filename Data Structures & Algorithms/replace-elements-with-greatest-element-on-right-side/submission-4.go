func replaceElements(arr []int) []int {
	mx := -1
	for i := len(arr)-1; i > -1; i-- {
		temp := arr[i]
		arr[i] = mx
		mx = max(temp, mx)
	}
	return arr
}
