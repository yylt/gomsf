package gomsf

import (
	"github.com/vmihailenco/msgpack"

	"fmt"
	"strings"
)

const (
	cslcreate = "console.create"
	csllist = "console.list"
	cslrm = "console.destroy"
	cslread = "console.read"
	cslwrite= "console.write"
	csltabs = "console.tabs"
	cslsskill = "console.session_kill"
	cslssdetach = "console.session_detach"
)

type ConsoleInfo struct {
	Id string `msgpack:"id"`
	Prompt string `msgpack:"prompt"`
	Busy bool `msgpack:"busy"`
}

type ConsoleDataInfo struct {
	Data string `msgpack:"data"`
	Prompt string `msgpack:"prompt"`
	Busy bool `msgpack:"busy"`
}

func (m *MsfGo) ConsoleCreate() (*ConsoleInfo,error) {
	var (
		bs []byte
		err error
		token = m.GetToken()
	)
	if token==""{
		return nil,ErrNotAuth
	}
	bodys := []string{cslcreate,token}
	bs,err=msgpack.Marshal(bodys)
	if err!= nil{
		return nil,err
	}
	bs,err=m.send(false,bs)
	if err!= nil{
		return nil,err
	}
	var ct ConsoleInfo
	err=msgpack.Unmarshal(bs,&ct)
	if err!= nil{
		return nil,err
	}
	return &ct,nil
}

func (m *MsfGo) ConsoleList() ([]*ConsoleInfo,error) {
	var (
		bs []byte
		err error
		token = m.GetToken()
		infos []*ConsoleInfo
	)
	if token==""{
		return nil,ErrNotAuth
	}
	bodys := []string{csllist,token}
	bs,err=msgpack.Marshal(bodys)
	if err!= nil{
		return nil,err
	}
	bs,err=m.send(false,bs)
	if err!= nil{
		return nil,err
	}
	var ct = map[string]*ConsoleInfo{}
	err=msgpack.Unmarshal(bs,&ct)
	if err!= nil{
		return nil,err
	}
	for _,v:=range ct{
		infos=append(infos,v)
	}
	return infos,nil
}

func (m *MsfGo) ConsoleDestroy(id string) (error) {
	var (
		bs []byte
		err error
		token = m.GetToken()
	)
	if token==""{
		return ErrNotAuth
	}
	bodys := []string{cslrm,token,id}
	bs,err=msgpack.Marshal(bodys)
	if err!= nil{
		return err
	}
	_,err=m.send(true,bs)
	return err
}

func (m *MsfGo) ConsoleWrite(cmd string) (error) {
	var (
		bs []byte
		err error
		token = m.GetToken()
	)
	if token==""{
		return ErrNotAuth
	}
	cmd = fmt.Sprintf("%s\n",cmd)
	length := len(cmd)
	bodys := []string{cslwrite,token,cmd}
	bs,err=msgpack.Marshal(bodys)
	if err!= nil{
		return err
	}
	var ct = map[string]int{}
	bs,err=m.send(false,bs)
	err=msgpack.Unmarshal(bs,&ct)
	if err!= nil{
		return err
	}
	if v,ok:=ct["wrote"];ok{
		if v==length{
			return nil
		}
		return fmt.Errorf("wrong length %d",v)
	}
	return fmt.Errorf("wrong return-type: %v",ct)
}

func (m *MsfGo) ConsoleRead(id string) (string,error) {
	var (
		bs []byte
		err error
		token = m.GetToken()
	)
	if token==""{
		return "",ErrNotAuth
	}
	bodys := []string{cslread,token,id}
	bs,err=msgpack.Marshal(bodys)
	if err!= nil{
		return "",err
	}
	var ct ConsoleDataInfo
	bs,err=m.send(false,bs)
	err=msgpack.Unmarshal(bs,&ct)
	if err!= nil{
		return "",err
	}
	return strings.TrimSpace(ct.Data),nil
}
