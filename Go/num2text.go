package main

import ("fmt"
        "math")

/*Scrivere un programma num2text.go per convertire un numero
intero non negativo nella sequenza delle parole corrispondenti
alle sue cifre.
Il programma legge un intero non negativo da standard input,
per ogni nuova (non incontrata finora) cifra del numero
chiede il nome corrispondente (e alimenta un dizionario), e infine
stampa la sequenza delle parole corrispondenti alle sue cifre.
Ad esempio, per il numero 203, il programma stampa
due - zero - tre*/




func main()  {
  var dizionario map[int]string
  var numero int
  var c int
  dizionario = make(map[int]string)

  fmt.Scan(&numero)
  num2 := numero

  //aggiorno il dizionario se le cifre sono presenti
  for numero > 0{
    cifra := numero % 10
    _, ok := dizionario[cifra]
    if ok == false{
      fmt.Print("Parola per ",cifra," ? ")

      var s string
      fmt.Scan(&s)
      dizionario[cifra] = s
    }
    c++
    numero /= 10
  }
  //fmt.Println(dizionario)

  //stampo la cifra a parole
  for num2 > 0{
    div := int(math.Pow10(c-1))
    cifra := num2 / div

    //fmt.Println(stamp)
    fmt.Print(dizionario[cifra]," - ")
    num2 = num2 % div
    c--

  }

}
