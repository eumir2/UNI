package main

import "fmt"
import "os"
import "strconv"

/*che emetta nel flusso d'uscita una clessidra di base 2*n-1*/
func clessidra(n, offset int) {
  //caso base
  if n == 0{
    return
  }
  //clessidra(n-1,offset+1)

  if n == 1{
    for i := 0; i < offset; i++{
      fmt.Print(" ")
    }
    fmt.Print("#")
    fmt.Println()
  }else{

    //stampo gli spazi
    for i := 0; i < offset; i++{
      fmt.Print(" ")
    }
    //stampo gli asterischi
    for i := 0; i < 2*n-1; i++{
      fmt.Print("#")
    }
    //fmt.Printf("\tn:%d o:%d",n,offset)
    fmt.Println()

    clessidra(n-1,offset+1)

    for i := 0; i < offset; i++{
      fmt.Print(" ")
    }
    //stampo gli asterischi
    for i := 0; i < 2*n-1; i++{
      fmt.Print("#")
    }
    //fmt.Printf("\tn:%d o:%d",n,offset)
    fmt.Println()

    /*fmt.Printf("%3d%3d",n,offset)
    fmt.Println()*/
  }
}


func main()  {
  n,_ := strconv.Atoi(os.Args[1])

  //fmt.Println(n)

  clessidra(n,0)
}
