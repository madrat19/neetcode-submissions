func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	var m int
	for l <= r {
		m = (l + r) / 2
		if matrix[m][0] <= target && matrix[m][len(matrix[m])-1] >= target {
			break
		} else if matrix[m][0] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	l, r = 0, len(matrix[m])-1
	for l <= r {
		mid := (l + r) / 2
		if matrix[m][mid] == target {
			return true
		} else if matrix[m][mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
