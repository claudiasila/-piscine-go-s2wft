package main

import (
	"fmt"
	piscine "piscine/quest-04"
)

func main() {
	fmt.Printf("%c", piscine.FirstRune("Hello!"))
	fmt.Printf("%c", piscine.FirstRune("Salut!"))
	fmt.Printf("%c", piscine.FirstRune("Ola!"))
	fmt.Println()
}
