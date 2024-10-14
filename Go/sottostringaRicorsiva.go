package main

import "fmt"

func sottoStringaRicorsiva(s1,s2 string)bool{
/*  fmt.Println("\n------------\n",s1)
  fmt.Println(s2,"\n----------\n")*/
  //fmt.Println(len(s2))

  if len(s1) < len(s2){
    return false
  }
  if len(s2) == 0{
    //fmt.Println("qui")
    return true
  }
  if s1[0] == s2[0]{
    return sottoStringaRicorsiva(s1[1:],s2[1:])
  }else{
    return sottoStringaRicorsiva(s1[1:],s2)
  }
}


func main(){
  s1 := "ewghbkdfweitgdfciewnciuergiuciaoejkfbierubuiebe"
  s2 := "ciao"

  fmt.Println(sottoStringaRicorsiva(s1,s2))
}
