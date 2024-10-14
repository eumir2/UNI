package main

import (
  "fmt"
  "os"
  "bufio"
)

/*Scrivere un programma (stampaInversa.go) che, dato un testo in input,
lo stampa invertendo l'ordine delle parole (stampando l'ultima per prima,
la penultima per seconda, ecc).
Leggere una stringa alla volta, utilizzando il metodo Split con la funzione di split bufio.ScanWords*/

func main()  {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  var s []string

  for scanner.Scan(){
    s = append(s,scanner.Text())
  }

  for i := len(s)-1; i >= 0; i--{
    fmt.Print(s[i]," ")
  }
}
