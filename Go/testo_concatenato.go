package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var testoLetto string
	testoLetto = LeggiTesto()

	fmt.Println("Leggi testo")
	fmt.Println(testoLetto)

}

func LeggiTesto() (testo string) {
	fmt.Println("Inserisci del testo")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}
