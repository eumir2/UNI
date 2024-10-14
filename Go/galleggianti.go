package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* Prende una matrice (slice bidimensionale) dove m[r][c] è l'elemento
   sulla r-esima riga (0 è la riga più in alto) e la c-esima colonna (0
   è la colonna più a sinistra), e la modifica portando gli * in alto in
   ciascuna colonna */
func galleggia(m [][]string, nr int, nc int) {
	for ultima := nr; ultima > nr-2; ultima-- { // Sto per sistemare la riga #ultima - 1
		for c := 0; c < nc; c++ {
			for r := 0; r < ultima-1; r++ {
				if m[r+1][c] == "*" && m[r][c] != "*" {
					m[r+1][c], m[r][c] = m[r][c], m[r+1][c]
				}
			}
		}
	}
}

func readInput(nr int, nc int) [][]string {
	risultato := [][]string{} // oppure: var risultato [][]string
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < nr; i++ {
		scanner.Scan()                                          // Mi preparo a leggere una riga
		riga := scanner.Text()                                  // Leggo una riga
		risultato = append(risultato, strings.Split(riga, " ")) // Separo la riga usando gli spazi e accodo al risultato
	}
	return risultato
}

func printOutput(m [][]string, nr int, nc int) {
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
}

func main() {
	nr, _ := strconv.Atoi(os.Args[1])
	nc, _ := strconv.Atoi(os.Args[2])
	m := readInput(nr, nc)
	galleggia(m, nr, nc)
	printOutput(m, nr, nc)
}
