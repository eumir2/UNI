package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	len := len(os.Args)

	if len == 1 {
		fmt.Println("nessun valore in ingresso")
		return
	}

	previous, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("elemento in posizione 1 non valido")
		return
	}

	for i := 2; i < len; i++ {
		follower, err := strconv.Atoi(os.Args[i])

		/*
			if err != nil {
				return
			}
		*/

		//controllo parità: diverse alternative
		//if (follower%2 == 0 && previous%2 == 0) || (follower%2 != 0 && previous%2 != 0) {
		if (follower+previous)%2 == 0 || err != nil {
			//bit := (previous) % 2
			//if follower%2 == bit { //in questo caso non servono due var distinte prev e follow
			//ma serve var bit in più
			fmt.Println("elemento in posizione", i, "non valido")
			return
		}
		previous = follower
		//bit = (bit + 1) % 2
	}
	fmt.Println("sequenza valida")
}
