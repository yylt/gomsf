package gomsf

import (
	"fmt"
	"errors"

	"github.com/vmihailenco/msgpack"
	"github.com/imroc/req"
)

type MsfGo struct {
	addr string
	token string
	ssl bool

	httpcli *req.Req
}

type result struct {
	Result string `msgpack:"result"`
}

var (
	defaultParm = req.Param{
		"Content-Type":"binary/message-pack",
		"Content-Length":1024,
	}
	msfdUrl string
	ErrNotAuth = errors.New("not auth")
	ErrAuthFailed = errors.New("auth failed")
)

func NewMsf(addr string,port int,ssl bool) *MsfGo{
	rq:=req.New()
	if ssl{
		rq.EnableInsecureTLS(true)
		msfdUrl = fmt.Sprintf("https://%s:%d/api/1.0/",addr,port)
	}else{
		msfdUrl = fmt.Sprintf("http://%s:%d/api/1.0/",addr,port)
	}

	return &MsfGo{
		addr: addr,
		ssl:ssl,
		httpcli:rq,
	}
}

func (m *MsfGo) send(checkresult bool,cnt []byte) ([]byte,error){
	resq,err:=m.httpcli.Post(msfdUrl,defaultParm,cnt)
	if err!= nil {
		return nil,err
	}
	bs:=resq.Bytes()
	if checkresult{
		var res result
		msgpack.Unmarshal(bs,&res)
		if res.Result!= "success"{
			return nil,errors.New(res.Result)
		}
	}
	return bs,nil
}

func (m *MsfGo) GetToken() string {
	return m.token
}

func (m *MsfGo) SetToken(token string) {
	m.token=token
}