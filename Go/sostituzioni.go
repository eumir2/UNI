package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	frase := os.Args[1]
	vecchia := os.Args[2]
	nuova := os.Args[3]
	volte, _ := strconv.Atoi(os.Args[4])
	fmt.Println(frase)
	frase = sostituisci(frase, vecchia, nuova, volte)
	fmt.Println(frase)
}

func sostituisci(s, old, new string, n int) string {
	words := strings.Fields(s)
	var count int
	var news string
	for _, x := range words {
		if x == old {
			count++
		}
		if count == n-1 {
			x = new
			count++
		}
		news += x + " "
	}
	return news
}
