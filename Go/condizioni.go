package main
import "fmt"

/*Scrivere un programma Go condizioni.go che testa, una a una, le seguenti condizioni.
Implementarne una alla volta, testarla su almeno due input (uno che la verifica
e uno che la falsifica), e solo poi procedere alla successiva.
Il programma legge un valore da tastiera (del tipo indicato nella colonna “tipo”) e stampa true o
false, a seconda che la condizione sia verificata o no.*/

func main()  {
  var a,b,c,d,f,e,g,h,i,j,k,l int
  var x float64
/*
 fmt.Print("int uguale a 10: ")
  fmt.Scan(&a)

  if a == 10 {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }


  fmt.Print("int diverso da 10: ")
  fmt.Scan(&b)

  if b != 10 {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }


  fmt.Print("int diverso da 10 e da 20: ")
  fmt.Scan(&c)

  if c != 10 && c != 20 {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }

fmt.Print("int diverso da 10 o da 20: ")
fmt.Scan(&d)

if d != 10 || d != 20 {
  fmt.Println("true")
}else{
  fmt.Println("false")
}

  fmt.Print("int maggiore o uguale a 10: ")
  fmt.Scan(&e)

  if e >= 10 {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }


  fmt.Print("int compreso tra 10 e 20: ")
  fmt.Scan(&f)

  if f >= 10 && f <= 20 {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }

  fmt.Print("int compreso tra 10 e 20: ")
  fmt.Scan(&g)

  if g > 10 && g < 20 {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }


  fmt.Print("int compreso tra 10 e 20: ")
  fmt.Scan(&h)

  if h >= 10 && h < 20 {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }

  fmt.Print("int esterno all'intervallo 10 e 20: ")
  fmt.Scan(&i)

  if i < 10 || i > 20 {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }

  fmt.Print("int tra 10 e 20 o tra 30 e 40: ")
  fmt.Scan(&j)

  if (j >= 10 && j <= 20) || (j >= 30 && j <= 40)  {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }

  fmt.Print("int multiplo di 4 ma non di 100: ")
  fmt.Scan(&k)

  if k % 4 == 0 && k % 100 != 0 {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }

  fmt.Print("int dispari compreso tra 0 e 100: ")
  fmt.Scan(&l)

  if l % 2 != 0 && l >= 0 && l < 100 {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }

  fmt.Print("float64 vicino a 10.0 (non più lontano di 10^-6) (1e-6): ")
  fmt.Scan(&x)

  if x - 10.0 <= 1e-6 {
    fmt.Println("true")
  }else{
    fmt.Println("false")
  }

}
