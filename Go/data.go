package main

import (
	"fmt"
	"strconv"
)

type Data struct {
	gg uint
	mm uint
	aa uint
}

func CreaData(g uint, m uint, a uint) (d Data, ok bool) {
	if g > 0 && g < 32 && m > 0 && m < 12 && a > 0 {
		ok = true
		d.gg = g
		d.mm = m
		d.aa = a
	} else {
		ok = false
	}
	return d, ok
}

func aggiorna(d *Data) {
	if d.gg < 31 {
		d.gg = d.gg + 1
	} else {
		d.gg = 1
		d.mm += 1
	}
}

func palindroma(d Data) bool {
	var data string
	if d.gg < 10 {
		data = "0" + strconv.Itoa(int(d.gg))
	} else {
		data = strconv.Itoa(int(d.gg))
	}
	if d.mm < 10 {
		data += "0" + strconv.Itoa(int(d.gg))
	} else {
		data += strconv.Itoa(int(d.gg))
	}
	data += strconv.Itoa(int(d.aa))
	for i := 0; i < len(data)/2; i++ {
		if data[i] != data[len(data)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var (
		g, m, a uint
		d       Data
		ok      bool
	)
	g, m, a = 1, 2, 2020
	d, ok = CreaData(g, m, a)
	if ok {
		fmt.Println("creata:", d)
		aggiorna(&d)
		fmt.Println("data aggiornata:", d)
		if palindroma(d) {
			fmt.Println(d, "data palindroma")
		}
	}
}
