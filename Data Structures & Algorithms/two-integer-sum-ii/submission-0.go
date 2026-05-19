func twoSum(numbers []int, target int) []int {
	ans := []int{}
	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] == target {
			ans = append(ans, i+1, j+1)
			return ans
		}
		if numbers[i]+numbers[j] > target {
			j--
		}
		if numbers[i]+numbers[j] < target {
			i++
		}
	}
	return ans
}
