package gomsf

import (
	"github.com/vmihailenco/msgpack"
)

const (
	authlogin = "auth.login"
	authlogou = "auth.logout"
	tokenlist = "auth.token_list"
	tokenadd = "auth.token_add"
	tokengen= "auth.token_generate"
	tokenrm = "auth.token_remove"
)

type tokendt struct {
	Result string `msgpack:"result"`
	Token string `msgpack:"token"`
}

type tlistdt struct {
	Tokens []string `msgpack:"tokens"`
}

func (m *MsfGo) Login(user,pass string) (string,error) {
	bodys := []string{authlogin,user,pass}
	bs,err:=msgpack.Marshal(bodys)
	if err!= nil{
		return "",err
	}
	resp,err:=m.send(true,bs)
	if err!= nil{
		return "",err
	}

	var auth tokendt
	err = msgpack.Unmarshal(resp,&auth)
	if err!= nil{
		return "",err
	}
	return auth.Token,nil

}

func (m *MsfGo) Logout() error {
	var (
		token = m.GetToken()
	)
	if token==""{
		return ErrNotAuth
	}
	bodys := []string{authlogou,token}
	bs,err:=msgpack.Marshal(bodys)
	if err!= nil{
		return err
	}
	_,err=m.send(true,bs)
	return err
}


func (m *MsfGo) TokenAdd(newt string) error {
	var (
		token = m.GetToken()
	)
	if token==""{
		return ErrNotAuth
	}
	bodys := []string{tokenadd,newt}
	bs,err:=msgpack.Marshal(bodys)
	if err!= nil{
		return err
	}
	_,err=m.send(true,bs)
	return err
}

func (m *MsfGo) TokenGen( ) (string,error) {
	var (
		token = m.GetToken()
	)
	if token==""{
		return "",ErrNotAuth
	}
	bodys := []string{tokengen,token}
	bs,err:=msgpack.Marshal(bodys)
	if err!= nil{
		return "",err
	}
	var auth tokendt
	err = msgpack.Unmarshal(bs,&auth)
	if err!= nil{
		return "",err
	}
	return auth.Token,nil
}


func (m *MsfGo) TokenList() ([]string,error) {
	var (
		token = m.GetToken()
	)
	if token==""{
		return nil,ErrNotAuth
	}
	bodys := []string{tokenlist,token}
	bs,err:=msgpack.Marshal(bodys)
	if err!= nil{
		return nil,err
	}
	resp,err:=m.send(false,bs)
	if err!= nil{
		return nil,err
	}
	var t tlistdt
	err = msgpack.Unmarshal(resp,&t)
	if err!= nil{
		return nil,err
	}
	return t.Tokens,nil
}
