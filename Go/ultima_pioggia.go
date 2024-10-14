package main
import ("fmt"
        "bufio"
        "os"
        "strconv")

/*Scrivere un programma ultima_pioggia.go che legge da standard input una
sequenza di valori interi (terminata da EOF, ottenuto con ctrl-D)
che indicano i mm di pioggia caduti (0 se non ha piovuto) ogni giorno
in una sequenza successiva di giorni e stampa (l'indice del) l'ultimo
giorno in cui ha piovuto.*/

func main()  {

  index := 0
  i := 0
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
      l := scanner.Text()
      mm, _ := strconv.Atoi(string(l[0 : len(l)]))
      //fmt.Println(mm)
      if mm != 0 {
        index = i
      }
      i++
    }
    fmt.Println(index)
  }
