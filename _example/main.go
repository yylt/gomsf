package main

import (
  "fmt"
  "github.com/yylt/gomsf"
)

func main(){

  cli := gomsf.NewMsf("127.0.0.1:55553",true)
  token,err:=gomsf.Login(cli,"pass","pass")
  if err!= nil {
    fmt.Println(err)
    return
  }
  err = gomsf.Logout(cli,token.Token,token.Token)
  fmt.Printf("%v\n",err)

}
