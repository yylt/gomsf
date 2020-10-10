package main

import (
	"fmt"
	"log"

	"github.com/yylt/gomsf"

	"time"
)

func main() {

	cli := gomsf.NewMsf("127.0.0.1:55553", true,
		gomsf.WithDebug(true),
		gomsf.WithSkipVerify(true),
		gomsf.WithTimeOut(time.Second*5))

	token, err := gomsf.Login(cli, "pass", "pass")
	if err != nil {
		fmt.Println(err)
		return
	}

	exploits, err := gomsf.ModExploits(cli, token)
	log.Printf("exploit %v, err %v", exploits, err)

	err = gomsf.Logout(cli, token, token)
	fmt.Printf("%v\n", err)

}
