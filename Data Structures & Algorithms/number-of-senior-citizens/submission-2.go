func countSeniors(details []string) int {
	ans := 0
    for _, str := range details {
		age := int(str[11] - '0') * 10 + int(str[12] - '0')
		if age > 60 {
			ans++
		}
	}
	return ans
}