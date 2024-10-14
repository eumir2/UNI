/*Nel linguaggio farfallino ciascuna vocale non accentata ( vocale ) viene sostituita da una sequenza di tre
caratteri vocale-f-vocale . Per esempio, la vocale a viene sostituita dalla sequenza afa , la vocale e
dalla sequenza efe e così via. Se una vocale è maiuscola, anche la sequenza di tre caratteri che sostituisce
la vocale deve essere definita da caratteri maiuscoli (ad esempio, la vocale A viene sostituita dalla sequenza
AFA )*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var text string
	text = LeggiTesto()
	fmt.Println("Il testo in farfallino é", TraduciTestoInFarfallino(text))

}

func LeggiTesto() (text string) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Inserisci un testo su più righe: ")
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}
	return

}

func TraduciTestoInFarfallino(text string) (butter string) {

	for _, c := range text {
		butter += TraduciLetteraInFarfallino(c)
	}
	return

}

func TraduciLetteraInFarfallino(char rune) (text string) {

	switch char {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		text = string(char) + "f" + string(char)
	default:
		text = string(char)
	}
	return
}
