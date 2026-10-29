func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	ans[0] = []int{1}
	for i := 1; i < numRows; i++ {
		arr := make([]int, i+1)
			for j := 0; j <= i; j++ {
				index1, index2 := j-1, j
				a, b := 0, 0
				if index1 > -1 {
					a = ans[i-1][index1]
				}
				if index2 < i {
					b = ans[i-1][index2]
				}
				arr[j] = a + b
			} 
		ans[i] = arr
	}
	return ans
}
