type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return "п"
	}
	return strings.Join(strs, "р")
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "п" {
		return []string{}
	}
	return strings.Split(encoded, "р")
}
