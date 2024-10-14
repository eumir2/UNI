package main
import "fmt"
import "math/rand"
import "time"

/*Scrivere un programma indovina_numero.go che chiede all'utente di indovinare
un numero intero random x tra 1 e MAX, (dove MAX e` una costante definita nel programma)
e ripete la richiesta fino a che l'utente indovina, oppure fino a un massimo di
MAX/2 tentativi.
Il programma stampa il numero di tentativi che sono stati necessari per indovinare
oppure il messaggio "hai perso, il numero era x".
Se il numero digitato dall'utente è fuori dall'intervallo [1,MAX],
il tentativo non viene considerato e viene visualizzato il messaggio
"fuori intervallo!", senza interrompere l'esecuzione. */



func main()  {
  const MAX =  50
  c := 0 //flag per il controllo del numero
  //genero il numero random
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  numero := r1.Intn(MAX) + 1
  fmt.Println(numero)

  tmp := 0
  for i := 1; i < MAX / 2; i++{
    fmt.Scan(&tmp)

    if tmp < 1 || tmp > MAX {
      fmt.Println("fuori intervallo")
      i--
    }
    if tmp == numero{
      c = 1
      fmt.Println("Indovinato in",i,"tentativi")
      break
    }
  }
  if c == 0 {
    fmt.Println("hai perso, il numero era: ", numero)
  }
}
