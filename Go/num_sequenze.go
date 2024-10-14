package main
import ("fmt"
        "os"
        "bufio")

/*Scrivere un programma num_sequenze.go che legge da standard input una
sequenza di uni (1) e zeri (0) (terminata da un 2), che inizia e finisce con 1,
e stampa il numero di sottosequenze di zeri.
Ad esempio per input 1 1 0 0 1 0 1 0 0 0 1 1 1 0 1,
l'output è 4.
(si considerano anche quelle di lunghezza 1*/

func main()  {
  var c int
  var s string

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    l := scanner.Text()
    s = string(l)
    if s[len(s)-1] == '2'{
      break
    }
  }
  for i := 1; i < len(s) - 1; i++{
    if s[i] == '0'{
      c++
      if s[i+ 1] == '0'{
        i++
      }
    }
  }
  fmt.Println(c)
}
