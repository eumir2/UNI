package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var turni int
	var numGiocatori int
	var giocatori []string
	var score []int
	var scoreMax int
	var vincitoreTurno []string
	fmt.Print("Inserisci il numero dei giocatori ")
	fmt.Scan(&numGiocatori)
	fmt.Printf("Inserisci i nomi dei %d giocatori\n", numGiocatori)
	giocatori = make([]string, numGiocatori)
	for i := 0; i < numGiocatori; i++ {
		fmt.Scan(&giocatori[i])
	}
	score = make([]int, numGiocatori)
	fmt.Print("Inserisci il numero di turni ")
	fmt.Scan(&turni)
	vincitoreTurno = make([]string, turni)
	for i := 0; i < turni; i++ {
		fmt.Println("===TURNO ", i+1, "===")
		scoreMax = 0
		for j := 0; j < numGiocatori; j++ {
			rand.Seed(int64(time.Now().Nanosecond()))
			giocate := [2]int{0, 0}
			giocate[0] = rand.Intn(6)
			giocate[1] = rand.Intn(6)
			score[j] = giocate[0] + giocate[1]
			fmt.Print("          Giocatore ", giocatori[j], " lancia 2 dadi e fa: ", giocate[0], " ", giocate[1])
			if score[j] > scoreMax {
				scoreMax = score[j]
				vincitoreTurno[i] = giocatori[j]
			}
			fmt.Println("")
		}
		fmt.Println("Turno ", i+1, " vinto da ", vincitoreTurno[i])
		fmt.Println("")
	}
	fmt.Println("VITTORIE:")
	for m := 0; m < turni; m++ {
		fmt.Println("VINCITORE TURNO ", m+1, ": ", vincitoreTurno[m])
	}
}
