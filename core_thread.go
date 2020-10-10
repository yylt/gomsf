package gomsf

import (
	"fmt"
	"time"
)

const (
	coreThreadList = "core.thread_list"
	coreThreadKill = "core.thread_kill"

	coreSetg   = "core.setg"
	coreUnSetg = "core.unsetg"
)

type ThreadInfo struct {
	Status   string    `msgpack:"status"`
	Critical bool      `msgpack:"critical"`
	Name     string    `msgpack:"name"`
	Started  time.Time `msgpack:"started"`
}

//TODO Should remove options when failed?
func SetOption(cli MsfCli, token string, kvs ...string) error {
	var (
		err error
		sts = []string{coreSetg, token, "", ""}
	)
	if len(kvs) == 0 || len(kvs)%2 != 0 {
		return fmt.Errorf("kv list is not odd or nil!")
	}
	for i := 0; i < len(kvs); i += 2 {
		sts[2] = kvs[i]
		sts[3] = kvs[i+1]
		err = cli.Send(sts, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func UnSetOption(cli MsfCli, token string, keys ...string) error {
	var (
		err error
		sts = []string{coreUnSetg, token, ""}
	)
	if len(keys) == 0 {
		return fmt.Errorf("key list is nil!")
	}
	for i := 0; i < len(keys); i++ {
		sts[2] = keys[i]
		err = cli.Send(sts, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func ThreadList(cli MsfCli, token string) ([]*ThreadInfo, error) {
	var (
		err   error
		sts   = []string{coreThreadList, token}
		infos []*ThreadInfo
	)

	err = cli.Send(sts, &infos)
	return infos, err
}

func ThreadKill(cli MsfCli, token, id string) (*Generic, error) {
	var (
		sts = []string{coreThreadKill, token, id}
		gen *Generic
	)

	err := cli.Send(sts, &gen)
	return gen, err
}
