package main

import ("fmt"
        "bufio"
        "os")

/*Scrivi un programma min_lunghezza.go che legge da standard input una serie di stringhe, terminata da EOF (da tastiera si ottiene digitando ctrl D), e stampa la stringa più corta tra quelle lette.
Se ce ne fossero più di una della lunghezza minima, stampa la prima di queste.
Cmpila e Caricalo poi su upload*/


func main()  {

  var sMin string
  var lMin  int = 100   //do per scontato che non si inseriscano stringhe più lunghe di 100 caratteri

  scanner :=  bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
      l := scanner.Text()
      s := string(l)

      if len(s) < lMin{
        sMin = s
        lMin = len(s)
      }
  }
  fmt.Println(sMin)
}
