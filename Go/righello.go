package main

import "fmt"
import "os"
import "strconv"

func righello(n int)  {
  if n == 0{
    return
  }else{
    righello(n-1)
    fmt.Println()
    for i := 0; i < n; i++{
      fmt.Print("-")
    }
    fmt.Println()
    righello(n-1)
  }
}

func main()  {
  n , _ := strconv.Atoi(os.Args[1])
  fmt.Println(n)

  righello(n)
}
