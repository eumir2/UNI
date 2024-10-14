package main

import "fmt"

func collatz(n int)int{
  if n == 0{
    fmt.Println(0)
  }
  if n == 1{
    fmt.Println(1)
    return 1
  }
  fmt.Println(n)
  if n % 2 == 0{
    return  collatz(n/2)
  }else{
    return  collatz(n * 3 + 1)
  }
}

func main()  {
  var n int

  fmt.Scan(&n)

  collatz(n)
}
