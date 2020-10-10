package gomsf

import "fmt"

const (
	cslcreate   = "console.create"
	csllist     = "console.list"
	cslrm       = "console.destroy"
	cslread     = "console.read"
	cslwrite    = "console.write"
	csltabs     = "console.tabs"
	cslsskill   = "console.session_kill"
	cslssdetach = "console.session_detach"
)

type ConsoleInfo struct {
	Id     string `msgpack:"id"`
	Prompt string `msgpack:"prompt"`
	Busy   bool   `msgpack:"busy"`
}

type ConsoleDataInfo struct {
	Data   string `msgpack:"data"`
	Prompt string `msgpack:"prompt"`
	Busy   bool   `msgpack:"busy"`
}

type ConsoleTabInfo struct {
	Tabs []string `msgpack:"tabs"`
}

func ConsoleCreate(cli MsfCli, token string) (*ConsoleInfo, error) {
	var (
		retv *ConsoleInfo
		sts  = []string{cslcreate, token}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func ConsoleDestory(cli MsfCli, token, consoleId string) (*Generic, error) {
	var (
		retv *Generic
		sts  = []string{cslrm, token, consoleId}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func ConsoleList(cli MsfCli, token string) ([]*ConsoleInfo, error) {
	var (
		retv []*ConsoleInfo
		sts  = []string{csllist, token}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func ConsoleWrite(cli MsfCli, token, consoleId string, data []byte) error {
	var (
		sts = []string{cslwrite, token, consoleId, string(data)}
	)
	var ret struct {
		Wrote int `msgpack:"wrote"`
	}
	var r = &ret
	err := cli.Send(sts, &r)
	if err != nil {
		return err
	}
	if ret.Wrote != len(data) {
		return fmt.Errorf("write %d, but expect %d", ret.Wrote, len(data))
	}
	return nil
}

func ConsoleRead(cli MsfCli, token, consoleId string) (*ConsoleDataInfo, error) {
	var (
		retv *ConsoleDataInfo
		sts  = []string{cslread, token, consoleId}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func ConsoleSessDetach(cli MsfCli, token, consoleId string) (*Generic, error) {
	var (
		retv *Generic
		sts  = []string{cslssdetach, token, consoleId}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func ConsoleSessKill(cli MsfCli, token, consoleId string) (*Generic, error) {
	var (
		retv *Generic
		sts  = []string{cslsskill, token, consoleId}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func ConsoleTabs(cli MsfCli, token, consoleId, input string) (*ConsoleTabInfo, error) {
	var (
		retv *ConsoleTabInfo
		sts  = []string{csltabs, token, consoleId, input}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}
