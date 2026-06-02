package piscine

func Capitalize(s string) string {
	result := ""
	newWord := true

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = r + 32
		}

		if newWord && r >= 'a' && r <= 'z' {
			r = r - 32
		}

		result += string(r)

		if (r >= 'A' && r <= 'Z') ||
			(r >= 'a' && r <= 'z') ||
			(r >= '0' && r <= '9') {
			newWord = false
		} else {
			newWord = true
		}
	}

	return result
}
