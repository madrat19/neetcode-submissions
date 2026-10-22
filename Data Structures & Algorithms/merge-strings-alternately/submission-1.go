
func mergeAlternately(word1 string, word2 string) string {
	l1, l2 := len(word1), len(word2)
	builder := strings.Builder{}
	builder.Grow(l1+l2)
	i, j := 0, 0
	for i < l1 && j < l2 {
		builder.WriteByte(word1[i])
		builder.WriteByte(word2[j])
		i++
		j++
	}
	if i == l1 {
		builder.WriteString(word2[j:])
	}
	if j == l2 {
		builder.WriteString(word1[i:])
	}
	return builder.String()
}
