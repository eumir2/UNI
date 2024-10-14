package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BottigliaVino struct {
	nome     string
	anno     int
	grado    float32
	quantita int
}

func (b BottigliaVino) String() string {
	return b.nome + "," + strconv.Itoa(b.anno) + "," + strconv.FormatFloat(float64(b.grado), 'f', -1, 32) + "," + strconv.Itoa(b.quantita)
}

func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool) {
	ok := true
	bottiglia := BottigliaVino{}

	if nome == "" || anno <= 0 || quantita <= 0 || grado <= 0 {
		ok = false
	} else {
		bottiglia.nome = nome
		bottiglia.anno = anno
		bottiglia.grado = grado
		bottiglia.quantita = quantita
	}

	return bottiglia, ok

}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	bottiglia := BottigliaVino{}
	ok := false
	parametri := strings.Split(riga, ",")
	if len(parametri) == 4 {
		anno, _ := strconv.Atoi(parametri[1])
		grado, _ := (strconv.ParseFloat(parametri[2], 32))
		quantita, _ := strconv.Atoi(parametri[3])
		bottiglia, ok = CreaBottiglia(parametri[0], anno, float32(grado), quantita)
	}
	return bottiglia, ok
}

func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino) {
	*cantina = append(*cantina, bot)
}

func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino) {
	bottiglia, ok := CreaBottigliaDaRiga(bot)
	if ok {
		AggiungiBottiglia(bottiglia, cantina)
	}

}

func ContaPerTipo(nome string, cantina []BottigliaVino) int {
	quantita := 0
	for _, v := range cantina {
		if v.nome == nome {
			quantita++
		}
	}
	return quantita
}

func main() {
	nomeFile := os.Args[1]
	//var vini []BottigliaVino
	file, _ := os.Open("./" + nomeFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bottiglia, ok := CreaBottigliaDaRiga(line)
		if ok {
			fmt.Println(bottiglia.String())
		}
	}
}
