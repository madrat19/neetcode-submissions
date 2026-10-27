func countSeniors(details []string) int {
	ans := 0
    for _, str := range details {
		age, _ := strconv.Atoi(str[11:13])
		if age > 60 {
			ans++
		}
	}
	return ans
}