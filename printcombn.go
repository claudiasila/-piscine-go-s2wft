package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}

	for {
		printDigits(arr, n)

		if arr[0] == 10-n {
			break
		}

		i := n - 1

		for arr[i] == 9-n+i {
			i--
		}

		arr[i]++

		for j := i + 1; j < n; j++ {
			arr[j] = arr[j-1] + 1
		}
	}

	z01.PrintRune('\n')
}

func printDigits(arr []int, n int) {
	for i := 0; i < n; i++ {
		z01.PrintRune(rune(arr[i]) + '0')
	}

	if arr[0] != 10-n {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}
