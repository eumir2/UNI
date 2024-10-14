package main

import (
	"fmt"
	"os"
)

func main() {
	var words string
	phrase := os.Args[1:]
	for _, x := range phrase {
		words += x + " "
	}
	if eterogramma(words) {
		fmt.Println("E' un eterogramma")
	} else {
		fmt.Println("Non è un eterogramma")
	}
}

func eterogramma(word string) bool {
	var check bool
	runes1 := []rune(word)
	var runes2 []rune
	for _, z := range runes1 {
		if z != 32 {
			runes2 = append(runes2, z)
		}
	}
	for i, x := range runes2 {
		for j, y := range runes2 {
			if x == y && i != j {
				check = false
				break
			} else {
				check = true
			}
		}
	}
	return check
}
