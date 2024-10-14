package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	WordScan()
}
func WordScan() {
	var line, word string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Inserisci una parola")
	fmt.Print("\t")
	for scanner.Scan() {
		word = ""
		line = scanner.Text()
		if line == strconv.Itoa(0) {
			break
		}
		for _, x := range line {
			word += string(x)
			if string(x) == "" || string(x) == " " {
				break
			}
		}
		fmt.Println(word)
		ParseWord(word)
		fmt.Println("\nInserisci una parola")
		fmt.Print("\t")
	}

}

func ParseWord(word string) {
	var vowel rune
	var ok bool
	var err bool
	var count int
	for _, character := range word {
		character = unicode.ToLower(character)
		if character < 'a' || character > 'z' {
			err = false
			break
		}
		if character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u' {
			err = true
			if character == vowel {
				ok = true
				count++
			} else if count == 0 {

				vowel = character
				count++
			} else if character != vowel && string(vowel) != "" {
				ok = false
				break
			}
		}
	}
	if ok == true && err == true {
		fmt.Printf("\ncorretto %c %d", vowel, count)
	} else if ok == false && err == true {
		fmt.Println("sbagliato")
	} else if err == false {
		fmt.Println("errore")
	}
}
