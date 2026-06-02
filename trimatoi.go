package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	found := false

	for _, r := range s {
		if r == '-' && !found {
			sign = -1
		}

		if r >= '0' && r <= '9' {
			result = result*10 + int(r-'0')
			found = true
		}
	}

	return result * sign
}
