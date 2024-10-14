package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var testo string
	fmt.Print("Inserisci un testo da invertire (TERMINA CON CTRL+D/CTRL+Z)\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	invertiStringa(testo)
}

func invertiStringa(s string) {
	for i := len(s); i > 0; i-- {
		fmt.Printf("%c", s[i-1])
	}
}
