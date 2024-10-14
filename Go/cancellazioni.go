package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cancella(lista []string) []string {
	var j int
	var err error
	var nuovaLista []string

	for i := 0; i < len(lista); i++ {
		if j > 0 {
			lista[i] = ""
			j--
		} else {
			j, err = strconv.Atoi(lista[i])
			if err == nil {
				lista[i] = ""
			} else {
				j = 0
			}
		}
	}

	for _, val := range lista {
		if val != "" {
			nuovaLista = append(nuovaLista, val)
		}
	}

	return nuovaLista
}

func main() {

	var testoFile string

	if len(os.Args) <= 1 {
		fmt.Printf("Fornire 1 nome di file\n")
		return
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("File non accessibile\n")
		return
	}
	defer f.Close()

	b := bufio.NewScanner(f)

	for b.Scan() {
		testoFile += b.Text() + " "
	}
	listaIniziale := strings.Split(strings.TrimSpace(testoFile), " ")

	fmt.Println(listaIniziale)

	lista := cancella(listaIniziale)

	fmt.Println(lista)
}
