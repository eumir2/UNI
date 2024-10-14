package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	var phrase string
	var count int
	for i := 1; i < len(os.Args); i++ {
		word := os.Args[i]
		phrase += TrasformaParola(word, count) + " "
		count++
	}
	fmt.Println(phrase)
}

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {
	chars := []rune(parola)
	if posizione%2 == 0 {
		for i, char := range chars {
			if i%2 == 0 {
				parolaTrasformata += string(unicode.ToUpper(char))
			} else {
				parolaTrasformata += string(char)
			}
		}
	} else {
		for i, char := range chars {
			if i%2 == 0 {
				parolaTrasformata += string(char)
			} else {
				parolaTrasformata += string(unicode.ToUpper(char))
			}
		}
	}
	return parolaTrasformata
}
