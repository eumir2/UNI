package main

import "fmt"

type Point3 struct {
	x float64
	y float64
	z float64
}

func main() {
	p1 := Point3{1.3, 2.5, -2}
	p2 := Point3{-3, -2.5, -2}
	proiezione(&p1)
	fmt.Println(p1)
	s := []Point3{p1, p2, {}}
	fmt.Println("punti con z 0:", punti2(s))
}

func proiezione(point *Point3) {
	point.z = 0
}

func punti2(s []Point3) []Point3 {
	var s2 []Point3
	for _, x := range s {
		if x.z == 0 {
			s2 = append(s2, x)
		}
	}
	return s2
}
