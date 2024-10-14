package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var testo string
	var testoTrasformato string
	testo = LeggiTesto()
	testoTrasformato = Spazia(testo)
	fmt.Printf("Risultato:\n%s", testoTrasformato)
}

func LeggiTesto() (text string) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Inserisci un testo su più righe: \n")
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}
	return

}

func Spazia(s string) string {
	var frase string
	for _, x := range s {
		if x == '\n' {
			frase += "\n"
		}
		if unicode.IsSpace(x) {
		} else {
			frase += string(x) + " "
		}
	}
	return frase
}
