package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	chiave := LeggiNumero()
	fmt.Println("Chiave di lettura: ", chiave)
	testo := LeggiTesto()
	fmt.Println("Testo prima della cifratura: ", testo)
	testoCifrato := CifraTesto(testo, chiave)
	fmt.Println("Testo dopo della cifratura: ", testoCifrato)
}

func LeggiTesto() string {
	var testo string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return testo
}

func LeggiNumero() int {
	var chiave int
	fmt.Println("Inserisci la chiave di lettura")
	fmt.Scan(&chiave)
	return chiave
}

func CifraCarattere(char rune, chiave int) rune {
	var somma rune
	if chiave > 0 {
		somma = char + rune(chiave)
	} else {
		chiave = -chiave
		somma = char - rune(chiave)
	}
	if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
		//se è una lettera dell'alfabeto inglese
		if unicode.IsUpper(char) {
			//se maiuscola
			switch {
			case somma > 'Z':
				r := somma - 'Z'
				char = 'A' + r - 1
				return char
			case somma < 'A':
				r := 'A' - somma
				char = 'Z' - r + 1
				return char
			case somma > 'A' && somma < 'Z':
				return somma
			}
		} else if unicode.IsLower(char) {
			//se minuscola
			switch {
			case somma > 'z':
				r := somma - 'z'
				char = 'a' + r - 1
				return char
			case somma < 'a':
				r := 'a' - somma
				char = 'z' - r + 1
				return char
			case somma > 'a' && somma < 'z':
				return somma
			}
		}
	} else {
		//se non è una lettera dell'alfabeto inglese
		if unicode.IsSpace(char) {
		} else {
			//fmt.Println(string(char), "non è una lettera dell’alfabeto inglese!")
		}
		return char
	}
	return char
}

func CifraTesto(s string, chiave int) string {
	var testoTrasformato string
	for _, x := range s {
		//cifro carattere per carattere
		y := CifraCarattere(x, chiave)
		testoTrasformato += string(y)
	}
	return testoTrasformato
}
