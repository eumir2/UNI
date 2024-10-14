package main

/*Scrivere un programma che:
legga da standard input un testo su più righe;
termini la lettura quando viene inserita da standard input una riga vuota ( "" );
ristampi il testo letto (riga vuota esclusa) sostituendo tutte le vocali con delle u .
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var testo string
	var testoTrasformato string
	fmt.Println("Inserisci un testo su più righe (termina con riga vuota):")
	testo = LeggiTesto()
	testoTrasformato = Garibaldi(testo)
	fmt.Printf("Risultato trasformazione \n%s", testoTrasformato)

}

func LeggiTesto() string {
	var frase string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		frase += scanner.Text() + "\n"
	}
	return frase
}

func TrasformaCarattere(c rune) rune {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		c = 'u'
	default:
	}
	return c
}

func Garibaldi(s string) string {
	var frase string
	for _, x := range s {
		x = TrasformaCarattere(x)
		frase += string(x)
	}
	return frase
}
