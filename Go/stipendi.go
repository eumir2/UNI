package main

import "fmt"

type Persona struct{
  nome,cognome string
  stipendio int
}

func stipendiDispari(p []Persona)int{
  conta := 0
  somma := 0
  for _,v := range p{
    if len(v.cognome) % 2 != 0{
      somma += v.stipendio
      conta++
    }
  }
  return somma / conta
}


func main(){
  p1 := Persona{"Mario","Dispari",15}
  p2 := Persona{"Luigi","Pari",12}
  p3 := Persona{"Waluigi","Dispari",15}
  p4 := Persona{"Wario","Dispari",15}
  p5 := Persona{"Toad","Pari",12}
  persone := []Persona{p1,p2,p3,p4,p5}
  fmt.Println(persone)
  fmt.Println(stipendiDispari(persone))
}
