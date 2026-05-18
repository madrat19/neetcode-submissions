func longestConsecutive(nums []int) int {
	l := len(nums)
	if l == 0 || l == 1 {
		return l
	}

	counter := make(map[int]struct{})
	for _, num := range nums {
		counter[num] = struct{}{}
	}

	starts := []int{}
	for num := range counter {
		if _, ok := counter[num-1]; !ok {
			starts = append(starts, num)
		}
	}

	ans := 1
	for _, start := range starts {
		cur := 1
		for next := start + 1; ; next++ {
			if _, ok := counter[next]; ok {
				cur++
			} else {
				break
			}
		}
		ans = max(ans, cur)
	}
	return ans
}
