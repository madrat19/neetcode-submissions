func validPalindrome(s string) bool {
	dLeft, dRight := false, false
	pLeft, pRight := true, true
	i, j := 0, len(s) - 1
	for i < j {
		if s[i] != s[j] {
			if dLeft {
				pLeft = false
				break
			} else {
				dLeft = true
				i++
			}
		} else {
			i++
			j--
		}
	}
	i, j = 0, len(s) - 1
	for i < j {
	if s[i] != s[j] {
		if dRight {
			pRight = false
			break
		} else {
			dRight = true
			j--
		}
	} else {
		i++
		j--
	}
}
	return pLeft || pRight
}



