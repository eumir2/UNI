package main

/*Scrivere un programma che legga da standard input una riga di testo e stampi in output il doppio del numero
nascosto all'interno della riga di testo, ovvero il doppio del numero che si ottiene concatenando le cifre
presenti all'interno della riga di testo. Il programma non stampa nulla se non è presente alcun numero
nascosto.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	frase := LeggiTesto()
	numero, err := NumeroNascosto(frase)
	if err != nil {
		fmt.Println("nessun numero nascosto")
	} else {
		fmt.Printf("Il doppio del numero nascosto è: %d (%d * 2)", numero*2, numero)
	}
}

func LeggiTesto() string {
	var frase string
	fmt.Print("Inserisci la frase\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase = scanner.Text()
	return frase
}

func NumeroNascosto(testo string) (int, error) {
	var y string
	for _, x := range testo {
		if unicode.IsDigit(x) {
			y += string(x)
		} else {
		}
	}
	num, err := strconv.Atoi(y)
	return num, err
}
