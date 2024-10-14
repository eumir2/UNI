package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Persona struct {
	id      int
	nome    string
	cognome string
	età     int
}

type InsiemeTre struct {
	persona1 Persona
	persona2 Persona
	persona3 Persona
}

func main() {
	var p Persona
	var persone []Persona
	var links, notlinks []int

	file1 := os.Args[1]
	file2 := os.Args[2]

	if file1 == "Persone.txt" {
		f1, _ := os.Open("Persone.txt")
		scanner := bufio.NewScanner(f1)
		for scanner.Scan() {
			line := scanner.Text()
			lines := strings.Split(line, ";")
			p.id, _ = strconv.Atoi(lines[0])
			p.nome = lines[1]
			p.cognome = lines[2]
			p.età, _ = strconv.Atoi(lines[3])
			persone = append(persone, p)
		}
		defer f1.Close()
	}
	fmt.Println()
	if file2 == "Legami.txt" {
		f2, _ := os.Open("Legami.txt")
		scanner := bufio.NewScanner(f2)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			lines := strings.Split(line, ";")
			link1, _ := strconv.Atoi(lines[0])
			link2, _ := strconv.Atoi(lines[1])
			if len(links) == 0 {
				links = append(links, link1)
				if link1 != link2 {
					links = append(links, link2)
				}
			}
			notlinks = append(notlinks, link1)
			notlinks = append(notlinks, link2)
			ok := Appends(link1, links)
			if ok {
				links = append(links, link1)
			}
			ok = Appends(link2, links)
			if ok {
				links = append(links, link2)
			}
			defer f2.Close()
		}
		mp := LinksCompare(links, notlinks)
		fmt.Println(mp)
		media := MediaCombs(links)
		fmt.Println(media)
		Output(mp, persone, media)
	}
}

/*for _, p1 := range persone {
	if p1.età == links[len(links)-3] || p1.età == links[len(links)-2] || p1.età == links[len(links)-1] {
		fmt.Println(StringPersona(p1))
	}
}*/

func StringPersona(p Persona) string {
	return strconv.Itoa(p.id) + " - " + p.nome + " " + p.cognome + ": " + strconv.Itoa(p.età)
}

func Appends(link int, links []int) bool {
	for _, x := range links {
		if x == link {
			return false
		}
	}
	return true
}

func AppendAvg(link float64, links []float64) bool {
	for _, x := range links {
		if x == link {
			return false
		}
	}
	return true
}

func LinksCompare(sl, sl2 []int) map[int]int {
	var count int
	var counts []int
	for _, x := range sl {
		count = 0
		for _, y := range sl2 {
			if x == y {
				count++
			}
		}
		counts = append(counts, count)
	}

	var mp = make(map[int]int, len(sl))
	for i := 0; i < len(sl); i++ {
		mp[sl[i]] = counts[i]
	}
	return mp
}

func Reverse(s []float64) []float64 {
	var reverse []float64
	for i := len(s) - 1; i > -1; i-- {
		reverse = append(reverse, s[i])
	}
	return reverse
}

func MediaCombs(s []int) []float64 {
	var media []float64
	var num float64
	const k int = 3
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			for z := j + 1; z < len(s); z++ {
				num = (float64(s[i]) + float64(s[j]) + float64(s[z])) / float64(k)
				media = append(media, num)
			}
		}
	}
	sort.Float64s(media)
	media = Reverse(media)
	return media
}

func output(mp map[int]int, persone []Persona, s []float64) {
	for _, persona := range persone {

	}
}
