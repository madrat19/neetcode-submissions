
import (
	"slices"
)

func isAnagram(s string, t string) bool {
	runes1 := []rune(s)
	runes2 := []rune(t)
	slices.Sort(runes1)
	slices.Sort(runes2)
	s = string(runes1)
	t = string(runes2)
	return s == t
}
