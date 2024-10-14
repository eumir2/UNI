/*

Isogrammi
=========

Un isogramma è una parola o una frase senza caratteri che si ripetono, tranne spazi e trattini ('-') che possono apparire più volte.

Esempi di isogrammi sono:

	velocista
	tavoli
	conquista
	break-time
	dove si cena?
	come si fa?
	isogram
	six-pack

Non sono invece isogrammi:

	casa
	tavolo
	cruscotto
	questo no
	twelve-pack


Scrivere un programma in Go (il file deve chiamarsi 'isogrammi.go') dotato di:

- una funzione 'isIsogram(s string) bool'

	che restituisce true se la stringa s è un isogramma (ha tutti i caratteri diversi, tranne eventuali spazi e trattini ('-')

- una funzione main() che legge un numero arbitrario di righe di testo da standard input, per ciascuna riga  stabilisce se la riga stessa è un isogramma o no e stampa su standard output la riga seguita da SI/NO.

N.B. per segnalare al programma la fine dell'input, dare da tastiera invio seguito da ctrl D


Esempio di esecuzione
=====================

$ go run isogrammi.go
velocista
velocista: SI
break-time
break-time: NO
six-pack
six-pack: SI
hold-me-back
hold-me-back: SI
cruscotto
cruscotto: NO
isogram
isogram: SI

Nota Bene, se lo si lancia però dandogli un input redirezionato NON si vedono gli input, ad es.:
$ ./isogrammi < uno.input
velocista: SI
tavoli: SI
supercalifragilistichespiralidoso: NO
The Quick Brown Fox Jumps Over The Lazy Dog: NO
conquista: SI
break-time: NO
dove si cena?: NO
come si fa?: SI
isogram: SI
six-pack: SI
casa: NO
tavolo: NO
cruscotto: NO
diavolo: NO

*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
	//"strings"
	//"log"
)

func TestInput(t *testing.T) {
	lancia("uno.input", "uno.expected", t)
}

func lancia(inputF string, expectedF string, t *testing.T) {
	fmt.Println("### Questo test verifica output atteso!")

	subproc := exec.Command("./isogrammi") // presuppone che sia già stato compilato

	input, err := os.Open(inputF)
	subproc.Stdin = input

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

func TestNotIsogram(t *testing.T) {
	fmt.Println("-----non isogramma------------------------------")
	if isIsogram("superatissimo") {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}

func TestIsogram(t *testing.T) {
	fmt.Println("-----isogramma------------------------------")
	if !isIsogram("cavoli") {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}

func TestIsogramSpezzato(t *testing.T) {
	fmt.Println("-----isogramma spezzato------------------------------")
	if !isIsogram("cavoli due") {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}

func TestIsogramSpezzatoExtraChar(t *testing.T) {
	fmt.Println("-----isogramma spezzato extra char------------------------------")
	if !isIsogram("cavoli_ ! due # qp") {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}
