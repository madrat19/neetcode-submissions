func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n]++
	}
	buckets := make([][]int, len(nums)+1)
	for n, f := range counter {
		buckets[f] = append(buckets[f], n)
	}
	ans := []int{}
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, n := range buckets[i] {
			if k == 0 {
				return ans
			}
			ans = append(ans, n)
			k--
		}
	}
	return ans
}
