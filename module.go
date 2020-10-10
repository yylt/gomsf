package gomsf

import (
	"fmt"
	msgpack "github.com/vmihailenco/msgpack/v5"
)

type ModType string

const (
	ExploitType ModType = "exploit"
	AuxType     ModType = "auxiliary"
	PostType    ModType = "post"
	PayloadType ModType = "payload"
	EncoderType ModType = "encoder"
	NopType     ModType = "nop"
)

const (
	modexploit   = "module.exploits"
	modauxiliary = "module.auxiliary"
	modpost      = "module.post"
	modpayloads  = "module.payloads"
	modencoders  = "module.encoders"
	modnops      = "module.nops"

	modinfo                = "module.info"
	modoptions             = "module.options"
	modcompatible_payloads = "module.compatible_payloads"
	modcompatible_sessions = "module.compatible_sessions"

	modexecute = "module.execute"
)

type Modules struct {
	Module []string `msgpack:"modules"`
}

type ModuleDetail struct {
	Type     string   `msgpack:"type"`
	Desc     string   `msgpack:"description"`
	Fullname string   `msgpack:"fullname,omitempty"`
	Filepath string   `msgpack:"filepath"`
	Version  string   `msgpack:"version,omitempty"`
	Rank     string   `msgpack:"rank"`
	Authors  []string `msgpack:"authors"`
}

type ModOption struct {
	Name     string      `msgpack:"-"`
	Type     string      `msgpack:"type"`
	Required bool        `msgpack:"required"`
	Advanced bool        `msgpack:"advanced"`
	Evasion  bool        `msgpack:"evasion"`
	Desc     string      `msgpack:"desc"`
	Default  interface{} `msgpack:"default"`
	Enums    []string    `msgpack:"enums,omitempty"`
}

func ModExploits(cli MsfCli, token string) (*Modules, error) {
	var (
		retv *Modules
		sts  = []string{modexploit, token}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func ModAuxiliarys(cli MsfCli, token string) (*Modules, error) {
	var (
		retv *Modules
		sts  = []string{modauxiliary, token}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func ModPosts(cli MsfCli, token string) (*Modules, error) {
	var (
		retv *Modules
		sts  = []string{modpost, token}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func ModPayloads(cli MsfCli, token string) (*Modules, error) {
	var (
		retv *Modules
		sts  = []string{modpayloads, token}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func ModEncoders(cli MsfCli, token string) (*Modules, error) {
	var (
		retv *Modules
		sts  = []string{modencoders, token}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func ModNops(cli MsfCli, token string) (*Modules, error) {
	var (
		retv *Modules
		sts  = []string{modnops, token}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func InfoOfMod(cli MsfCli, token string, modtype ModType, modname string) (*ModuleDetail, error) {
	var (
		retv *ModuleDetail
		sts  = []string{modinfo, token, string(modtype), modname}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func OptionsOfMod(cli MsfCli, token string, modtype ModType, modname string) ([]*ModOption, error) {
	var (
		retv    map[string]*ModOption
		sts     = []string{modoptions, token, string(modtype), modname}
		options []*ModOption
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	for mod, value := range retv {
		value.Name = mod
		options = append(options, value)
	}
	return options, nil
}

func CompPayloadsMod(cli MsfCli, token string, exploitModName string) ([]string, error) {
	var (
		retv struct {
			Payloads []string `msgpack:"payloads"`
		}
		sts = []string{modcompatible_payloads, token, exploitModName}
	)
	var r = &retv
	err := cli.Send(sts, &r)
	if err != nil {
		return nil, err
	}
	return retv.Payloads, nil
}

func CompSessionsMod(cli MsfCli, token string, exploitModName string) ([]string, error) {
	var (
		retv struct {
			Sessions []string `msgpack:"sessions"`
		}
		sts = []string{modcompatible_sessions, token, exploitModName}
	)
	var r = &retv
	err := cli.Send(sts, &r)
	if err != nil {
		return nil, err
	}
	return retv.Sessions, nil
}

func Execute(cli MsfCli, token string, modType ModType, modName string, options map[string]interface{}) (int, error) {
	var (
		retv struct {
			Id int `msgpack:"job_id"`
		}
		sts = []string{modexecute, token, string(modType), modName}
	)
	switch modType {
	case ExploitType:
	case PostType:
	case AuxType:
	default:
		return 0, fmt.Errorf("Not support module type %s on execute operator", string(modType))
	}
	opts, err := msgpack.Marshal(options)
	if err != nil {
		return 0, err
	}
	sts = append(sts, string(opts))
	var r = &retv
	err = cli.Send(sts, &r)
	if err != nil {
		return 0, err
	}
	return retv.Id, nil
}
