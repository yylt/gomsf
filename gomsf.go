package gomsf

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/imroc/req"
	msgpack "github.com/vmihailenco/msgpack/v5"
)


type Generic struct {
	Restult string `msgpack:"result"`
}

type MsfGo struct {
	baseurl  string

	headers http.Header
	httpcli *req.Req
}

type MsfCli interface {
	Send([]string, interface{}) (error)
}

type Opt func(msfGo *MsfGo)

func WithDebug(debug bool) Opt{
	return func(msfGo *MsfGo) {
		if debug == true{
			req.Debug=true
			return
		}
		req.Debug=false
	}
}

func WithSkipVerify(ssl bool) Opt {
	return func(msfGo *MsfGo) {
		msfGo.httpcli.EnableInsecureTLS(ssl)
	}
}

func WithTimeOut(tm time.Duration) Opt {
	return func(msfGo *MsfGo) {
		msfGo.httpcli.SetTimeout(tm)
	}
}

func NewMsf(addr string, usessl bool,opts ...Opt) *MsfGo {
	var msfdUrl string
	rq := req.New()
	if usessl {
		msfdUrl = fmt.Sprintf("https://%s/api/1.0/", addr)
	} else {
		msfdUrl = fmt.Sprintf("http://%s/api/1.0/", addr)
	}

	msf := &MsfGo{
		baseurl:    msfdUrl,
		httpcli: rq,
		headers: http.Header{
			"Content-Type":   []string{"binary/message-pack"},
		},
	}

	for _,v :=range opts {
		v(msf)
	}
	return msf
}


func (m *MsfGo) Send(sts []string, save interface{})  (error) {
	var (
		buf = bufpool.Get().(*bytes.Buffer)
		err error
	)
	buf.Reset()
	defer bufpool.Put(buf)

	err = msgpack.NewEncoder(buf).Encode(sts)
	if err!= nil {
		return err
	}

	resq, err := m.httpcli.Post(m.baseurl, m.headers, buf.Bytes())
	if err != nil {
		return err
	}
	return msgpack.Unmarshal(resq.Bytes(), &save)
}
