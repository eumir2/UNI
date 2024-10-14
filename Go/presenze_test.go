/*
Presenze
========

Scrivere un programma in Go (il file deve chiamarsi 'presenze.go') dotato di

- una struttura Studente (nome, cognome, matr, presenze), dove:
	nome e cognome sono di tipo string
	matr e presenze di tipo int

- funzione incrementaPresenze(p_studente *Studente)
	che incrementa di 1 le presenze di studente

- una funzione main() che legge il nome di un file e un carattere da linea di comando

Il file conterrà una serie di righe, ciascuna con Nome, Cognome, Matr di uno studente, numero di presenze
(nomi e cognomi di una parola sola e con iniziali in A-Z, a-z)

Il programma deve creare una slice di studenti con i dati del file.

Il programma deve poi incrementare le presenze degli studenti della slice
il cui nome inizia con il carattere letto (secondo parametro linea di comando)
e produrre in output la lista degli studenti con la situazione delle presenze aggiornata.



Esempio
=======

Dato un file "uno.input" contenente:

Ida Unruh  381732 2
Carolynn Wimmer  93824923 6
Vada Furrow  28391 5
Maurine Even  28318 9
Lamar Tapp  3261725631 8
Sook Earle  763512 6
Denita Strothers  217632187 7
Carin Poteet  27651 7
Ina Blessing  29319812 5
Tish Billingsly  28738213 8
Felice Hermanson  5968598659 6
Everette Zeman  73627632 7
Agnus Birnbaum  29329832 6
Dania Crail  2329832 1
Violet Carnell  29387287 4
Marcelina Tuck  2736276 5
Freda Leonardo  2387872 7
Vina Clermont  832323 6
Gidget Phung  38723872 9
Ignacia Schaffner  23762372 4

L'esecuzione di:

$ go run presenze.go datiEsPresenze.input I

genererà:

Ida Unruh 381732 3
Carolynn Wimmer 93824923 6
Vada Furrow 28391 5
Maurine Even 28318 9
Lamar Tapp 3261725631 8
Sook Earle 763512 6
Denita Strothers 217632187 7
Carin Poteet 27651 7
Ina Blessing 29319812 6
Tish Billingsly 28738213 8
Felice Hermanson 5968598659 6
Everette Zeman 73627632 7
Agnus Birnbaum 29329832 6
Dania Crail 2329832 1
Violet Carnell 29387287 4
Marcelina Tuck 2736276 5
Freda Leonardo 2387872 7
Vina Clermont 832323 6
Gidget Phung 38723872 9
Ignacia Schaffner 23762372 5

*/

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"
	//"strings"
	//"log"
)

func TestInput(t *testing.T) {
	lancia("uno.input", "I", "uno.expected", t)
	lancia("uno.input", "C", "due.expected", t)
}

func lancia(inputF, lettera, expectedF string, t *testing.T) {
	fmt.Println("### Questo test verifica output atteso!")

	subproc := exec.Command("./presenze", inputF, lettera) // presuppone che sia già stato compilato

	//input, err := os.Open(inputF)
	//subproc.Stdin = input

	output, err := subproc.CombinedOutput()

	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	}

	//expected, err := os.Open(expectedF)
	expectedC, err := ioutil.ReadFile(expectedF)
	expectedS := string(expectedC)

	if string(output) != expectedS {
		t.Errorf("Output NON corrisponde all'atteso")
	}

	fmt.Printf("Output ATTESO:\n%s\n", expectedS)
	fmt.Printf("Output OTTENUTO:\n%s\n", string(output))

	subproc.Process.Kill()
}

func TestFunzione(t *testing.T) {
	fmt.Println("### Questo test verifica risultato atteso!")

	studente := Studente{"rossi", "mario", 2873612873, 56}

	incrementaPresenze(&studente)

	if studente.presenze != 57 {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}
