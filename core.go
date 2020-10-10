package gomsf

//TODO

const (
	coreAddModePath = "core.add_module_path"
	coreModeStat    = "core.module_stats"
	coreReloadMode  = "core.reload_modules"
	coreSave        = "core.save"
	coreVersion     = "core.version"
	coreStop        = "core.stop"
)

type ModuleInfo struct {
	Exploits string `msgpack:"exploits"`
	Auxs     string `msgpack:"auxiliary"`
	Posts    string `msgpack:"post"`
	Encoders string `msgpack:"encoders"`
	Nops     string `msgpack:"nops"`
	Payloads string `msgpack:"payloads"`
}

type VersionInfo struct {
	Version string `msgpack:"version"`
	Ruby    string `msgpack:"ruby"`
	Api     string `msgpack:"api"`
}

func AddModulePath(cli MsfCli, user, pass string) (*ModuleInfo, error) {
	var (
		retv *ModuleInfo
		sts  = []string{coreAddModePath, user, pass}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func Version(cli MsfCli, token string) (*VersionInfo, error) {
	var (
		retv *VersionInfo
		sts  = []string{coreVersion, token}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func Stop(cli MsfCli, token string) (*Generic, error) {
	var (
		retv *Generic
		sts  = []string{coreStop, token}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}
