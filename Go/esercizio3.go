package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	etichetta string
	ascissa   float64
	ordinata  float64
}

func main() {
	Punti := NuovoTragitto()
	lunghezza, punto := Lunghezza(Punti)
	stringa := String(punto)
	fmt.Println("Lunghezza percorso: ", lunghezza, "\nPunto oltre metà: ", stringa)
}

func NuovoTragitto() (tragitto []Punto) {
	var line string
	var punto Punto
	var coords []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			break
		}
		coords = strings.Split(line, ";")
		fmt.Println(coords)
		punto.etichetta = coords[0]
		punto.ascissa, _ = strconv.ParseFloat(coords[1], 64)
		punto.ordinata, _ = strconv.ParseFloat(coords[2], 64)
		tragitto = append(tragitto, punto)
	}
	return tragitto
}

func Distanza(p1, p2 Punto) float64 {
	return math.Pow(math.Pow(p1.ordinata-p2.ordinata, 2)+math.Pow(p1.ascissa-p2.ascissa, 2), 0.5)
}

func String(p Punto) string {
	return p.etichetta + " = (" + fmt.Sprint(p.ascissa) + "," + fmt.Sprint(p.ordinata) + ")"
}

func Lunghezza(tragitto []Punto) (float64, Punto) {
	var lunghezza float64
	var lunghezze []float64
	for i := 0; i < len(tragitto)-1; i++ {
		lunghezza += Distanza(tragitto[i], tragitto[i+1])
		lunghezze = append(lunghezze, lunghezza)
	}
	var count int
	for _, tratta := range lunghezze {
		if tratta > (lunghezza / 2) {
			return lunghezza, tragitto[count+1]
		} else {
			count++
		}
	}
	return lunghezza, tragitto[count]
}
