package piscine

// FirstRune returns the first rune of a string.
func FirstRune(s string) rune {
	for _, r := range s {
		return r
	}
	return 0
}
