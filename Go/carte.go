package main

import ("fmt"
        "math/rand"
        "time")

type Carta struct{
  valore,seme string
}

const SEMI = 4
const VALORI = 13
const CARTE = 52

func carta(n int)(Carta,bool)  {
  semi := [SEMI]string{"C","Q","F","P"}
  valori := [VALORI]string{"A","1","2","3","4","5","6","7","8","9","J","Q","K"}
  var c Carta

  s := n/VALORI
  v := n%VALORI
  //fmt.Println("valore: ",v)
  //fmt.Println("seme: ",s, semi[s])

  c.seme = semi[s]
  c.valore = valori[v]

  //fmt.Printf("%s %s\n",c.seme,c.valore)

  return c,true
}

func estraiCarta() Carta{
  val := rand.Intn(CARTE)
  fmt.Println(val)
  
  c,_ := carta(val)
  return c
}

func main(){
  rand.Seed(time.Now().UnixNano())

  c := estraiCarta()

  fmt.Println(c)

}
