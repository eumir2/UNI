/*
Luhn
====

Scrivere un programma in Go (il file deve chiamarsi 'luhn.go') che,
data su linea di comando una lista di lunghezza arbitraria di numeri di carta di credito,
ristampa ogni numero originale seguito da "valido" / "non valido" (virgolette escluse).

Se non sono presenti dati sulla linea di comando, il programma stampa "nessun input".

Per dati non numerici o valori con un numero di cifre diverso da 16, il programma stampa "dato non valido"

L'algoritmo di Luhn, che consente di controllare la validità di vari numeri, tra cui quelli delle carte di credito, si basa su tre passi:
- moltiplicare per due una cifra sì e una no, partendo dalla penultima cifra a destra e proseguendo verso sinistra.
- Dove la moltiplicazione ha dato un risultato a due cifre, sottrarre 9 dal prodotto
- Sommare tutte le cifre, sia quelle che si trovano in posizione pari, sia quelle che si trovano in posizione dispari
Se la somma complessiva è divisibile per 10 (la divisione non ha resto) il numero è valido.

Ad esempio, dato il numero (senza spazi)
4539 1488 0343 6467
si dovranno raddoppiare
4_3_ 1_8_ 0_4_ 6_6_

ottenendo:
8_6_ 2_7_ 0_8_ 3_3_

Sommando ora tutte le cifre
8569 2478 0383 3437

8+5+6+9+2+4+7+8+0+3+8+3+3+4+3+7 = 80

si ottiene un multiplo di 10, quindi questo numero è valido.


Esempi
=======

$ go run luhn.go 4539148803436467 8273123273520569 4024007190270651 4929876177122565 bastiancontrario 492987
4539148803436467 valido
8273123273520569 non valido
4024007190270651 valido
4929876177122565 non valido
bastiancontrario dato non valido
492987 dato non valido


$ go run luhn.go
nessun input

*/

package main

import (
	"fmt"
	"os/exec"
	"testing"
	//"strings"
	//"log"
)

func TestMainMultiplo(t *testing.T) {
	lancia(t, "4539148803436467 valido\n8273123273520569 non valido\n4024007190270651 valido\n4929876177122565 non valido\nbastiancontrario dato non valido\n492987 dato non valido\n", "4539148803436467", "8273123273520569", "4024007190270651", "4929876177122565", "bastiancontrario", "492987")
}

func TestMainNoInput(t *testing.T) {
	lancia(t, "no input\n")
}

func lancia(t *testing.T, expected string, in ...string) {
	fmt.Println("### Questo test verifica output atteso!")

	subproc := exec.Command("./luhn", in...) // presuppone che sia già stato compilato

	stdout, stderr := subproc.CombinedOutput()

	if stderr != nil {
		t.Errorf("Fallito: %s\n", stderr)
	}

	fmt.Printf("Actual output:\n%s\n", string(stdout))

	fmt.Printf("Expected output:\n%s\n", expected)

	if string(stdout) != expected {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}

	subproc.Process.Kill()
	//fmt.Println("### NON verifica correttezza!")
}
