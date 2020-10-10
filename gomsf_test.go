package gomsf

import (
	"time"
)

var (
	Cli   *MsfGo
	Token string

	defaultAddr   = "127.0.0.1:55553"
	defaultuser   = "user"
	defaultpasswd = "passwd"
)

func init() {
	Cli = NewMsf(defaultAddr, true,
		WithDebug(true),
		WithSkipVerify(true),
		WithTimeOut(time.Second*5))

	token, err := Login(Cli, defaultuser, defaultpasswd)
	if err != nil {
		panic(err)
		return
	}
	Token = token
}
