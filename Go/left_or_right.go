package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var value int
	var newvalue string
	var move string
	var count, win, lose int
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		value = randInt(1, 3)
		if value == 1 {
			newvalue = "s"
		} else {
			newvalue = "d"
		}
		count++
		fmt.Println("Dove sto guardando?")
		fmt.Scan(&move)
		if move == strconv.Itoa(0) {
			break
		}
		if move == newvalue {
			fmt.Println("Gisto hai vinto!")
			win++
		} else {
			fmt.Println("Sbagliato hai perso!")
			lose++
		}
	}
	fmt.Printf("Vinto: %.2f% Perso: %.2f%", float64(win)*100/float64(count), float64(lose)*100/float64(count))
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
