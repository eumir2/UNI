package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	num := LeggiNumero()
	mazzo := GeneraMazzo()
	EstraiCarte(mazzo, num)
}
func LeggiNumero() int {
	num, _ := strconv.Atoi(os.Args[1])
	return num
}

func GeneraMazzo() string {
	var mazzo string
	primo := '\U0001F0B1'
	for i := 0; i < 10; i++ {
		mazzo += string(primo)
		primo++
	}
	return mazzo
}
func EstraiCarta(mazzo string) (cartaEstratta rune, mazzoResiduo string) {
	mazzoRune := []rune(mazzo)
	randomIndex := rand.Intn(len(mazzoRune))
	for i, x := range mazzo {
		if i/4 == randomIndex {
			cartaEstratta = x
		}
		if x == cartaEstratta {
			mazzo = strings.ReplaceAll(mazzo, string(x), "")
		}
	}
	fmt.Printf("Estratta la carta %c", cartaEstratta)
	fmt.Printf("-Mazzo residuo %s\n", mazzo)
	return cartaEstratta, mazzo
}

func EstraiCarte(mazzo string, n int) {
	for i := 0; i < n; i++ {
		_, mazzo = EstraiCarta(mazzo)
	}
}
