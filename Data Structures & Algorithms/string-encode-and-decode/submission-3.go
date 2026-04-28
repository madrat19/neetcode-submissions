type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	result := ""
	for _, str := range strs {
		len := strconv.Itoa(len(str))
		result += len + "#" + str
	}
	return result
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		lenght, _ := strconv.Atoi(encoded[i:j])
		result = append(result, encoded[j+1:j+1+lenght])
		i = j + 1 + lenght
	}
	return result
}
