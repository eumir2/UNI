/*

Filtro
======

Scrivere un programma in Go (il file deve chiamarsi filtro.go) che
legge da standard input una stringa che contiene un solo simbolo della tastiera;
se il simbolo è L, T o Z, il programma stampa la lettere disegnata con asterischi (vedi sotto),
altrimenti stampa il messaggio "input non valido" (virgolette escluse).


Esempi
======

$ go run filtro.go
L
*
*
*
*
*****

$ go run filtro.go
T
*****
  *
  *
  *
  *

$ go run filtro.go
Z
*****
   *
  *
 *
*****

$ go run filtro.go
t
input non valido


*/

package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	//"log"
)

func TestMainMultiplo(t *testing.T) {
	lanciaEcontrolla("1", "input non valido\n", t)
	lanciaEcontrolla("ll", "input non valido\n", t)
	lanciaEcontrolla("L", "*\n*\n*\n*\n*****\n", t)
	lanciaEcontrolla("T", "*****\n  *\n  *\n  *\n  *\n", t)
	lanciaEcontrolla("Z", "*****\n   *\n  *\n *\n*****\n", t)
}

func lanciaEcontrolla(inp, outp string, t *testing.T) {
	fmt.Println("### Questo test verifica output atteso!")

	subproc := exec.Command("./filtro") // presuppone che sia già stato compilato

	subproc.Stdin = strings.NewReader(inp)
	stdout, stderr := subproc.CombinedOutput()

	if stderr != nil {
		t.Errorf("Fallito: %s\n", stderr)
	}

	fmt.Printf("Input:\n%s\n", inp)
	fmt.Printf("Output:\n%s\n", string(stdout))
	fmt.Printf("Expected output:\n%s\n", outp)

	if string(stdout) != outp {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}

	subproc.Process.Kill()
}
