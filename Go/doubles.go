package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var word string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\tInserisci una parola:")
	for scanner.Scan() {
		word = ""
		text := scanner.Text()
		if text == "0" {
			fmt.Println("\tFine")
			break
		}
		for _, x := range text {
			word += string(x)
			if string(x) == "" || string(x) == " " {
				break
			}
		}
		s := ParseWord(word)
		if s == nil {
			fmt.Println("\tNon ci sono doppie")
		} else {
			fmt.Println("\t", s)
		}
		fmt.Println("\nInserisci una parola")
		fmt.Print("\t")
	}
}

func ParseWord(word string) []string {
	var s []string
	var count int
	for i, x := range word {
		count = 0
		for j, y := range word {
			if x == y && i == j+1 {
				count++
			}
		}
		if count == 1 {
			s = append(s, string(x))
		}
	}

	return s
}
