package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var soglia int
	rand.Seed(int64(time.Now().Nanosecond()))
	soglia, _ = strconv.Atoi(os.Args[1])
	fmt.Println("Soglia:", soglia)
	generati, sopsoglia := Genera(soglia)
	fmt.Println("Valori generati: ", generati, "\nValori sopra soglia", sopsoglia)
}
func Genera(soglia int) ([]int, []int) {
	var numeroGenerato int
	var generati, soprasoglia []int

	for {
		numeroGenerato = rand.Intn(100)
		if numeroGenerato < soglia {
			generati = append(generati, numeroGenerato)
			break
		}
		generati = append(generati, numeroGenerato)
		if numeroGenerato > soglia {
			soprasoglia = append(soprasoglia, numeroGenerato)
		}
	}
	return generati, soprasoglia
}
