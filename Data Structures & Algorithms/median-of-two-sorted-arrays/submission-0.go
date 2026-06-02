func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n == 0 {
		return 0
	}
	i, j := 0, 0
	prev, cur := 0, 0
	for i+j <= n/2 {
		if i == len(nums1) {
			prev = cur
			cur = nums2[j]
			j++
		} else if j == len(nums2) {
			prev = cur
			cur = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			prev = cur
			cur = nums1[i]
			i++
		} else {
			prev = cur
			cur = nums2[j]
			j++
		}
	}
	if n%2 == 0 {
		return (float64(prev) + float64(cur)) / 2.0
	} else {
		return float64(cur)
	}
}
