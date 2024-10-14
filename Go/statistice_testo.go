// Stampare numero di parole e lunghezza media.

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	var parole, lunghezza int
	var totale string
	fmt.Print("Inserire del testo (TERMINARE CON CTRL+D): ")
	totale = LeggiTesto()
	parole, lunghezza = StatisticheParole(totale)
	fmt.Println("Hai inserito", parole, "parole di lunghezza media", float32(lunghezza)/float32(parole))

}

func LeggiTesto() (totale string) {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		totale += scanner.Text() + "\n"
	}
	return

}

func StatisticheParole(totale string) (numeroparole, lunghezza int) {

	var state bool = false
	for _, c := range totale {
		if unicode.IsLetter(c) {
			if state {
				lunghezza++
			} else {
				lunghezza++
				numeroparole++
				state = true
			}
		} else {
			if state {
				state = false
			}
		}
	}
	return
}
