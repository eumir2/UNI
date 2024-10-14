package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Studente struct {
	cognome, nome  string
	matr, presenze int
}

func incrementaPresenze(pStudente *Studente) {
	//fmt.Println(pStudente)
	//incrementa di 1 le presenze di studente
	//(*pStudente).presenze = (*pStudente).presenze + 1
	pStudente.presenze++
	//fmt.Println(pStudente)
}

func main() {
	//controllo args
	if len(os.Args) != 3 {
		fmt.Println("input non valido")
		return
	}
	nomeFile := os.Args[1]
	iniziale := os.Args[2][0]

	//apertura/controllo/chiusura file
	fileReader, err := os.Open(nomeFile)
	if err != nil {
		fmt.Println("errore apertura file ")
		return
	}
	defer fileReader.Close()

	var database []Studente

	//lettura file per righe
	scanner := bufio.NewScanner(fileReader)
	for scanner.Scan() {
		line := scanner.Text()

		//estrazione dati studente e creazione database
		// ATTENZIONE CHE NEL FILE NON DEVONO ESSERCI LE GRAFFE!
		dati := strings.Fields(line)
		matr, _ := strconv.Atoi(dati[2])
		pres, _ := strconv.Atoi(dati[3])
		studente := Studente{dati[0], dati[1], matr, pres}
		database = append(database, studente)
	}

	//aggiorna presenze e stampa dati
	for i := 0; i < len(database); i++ {
		if database[i].cognome[0] == iniziale {
			incrementaPresenze(&database[i])
		}
		s := fmt.Sprintf("%v", database[i])
		fmt.Println(s[1 : len(s)-1])
	}
}
