func merge(nums1 []int, m int, nums2 []int, n int) {
	ans := []int{}
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}
	ans = append(ans, nums1[i:m]...)
	ans = append(ans, nums2[j:]...)
	for i, e := range ans {
		nums1[i] = e
	}
}
