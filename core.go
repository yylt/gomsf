package gomsf
//TODO

const (
	coreAddModePath = "core.add_module_path"
	coreModeStat = "core.module_stats"
	coreReloadMode = "core.reload_modules"
	coreSave  = "core.save"
	coreVersion  = "core.version"
	coreStop  = "core.stop"
)


type ModuleInfo struct {
	Exploits  string `msgpack:"exploits"`
	Auxs  string `msgpack:"auxiliary"`
	Posts  string `msgpack:"post"`
	Encoders  string `msgpack:"encoders"`
	Nops  string `msgpack:"nops"`
	Payloads  string `msgpack:"payloads"`
}

func AddModulePath(cli MsfCli, user, pass string) (*ModuleInfo, error) {
	var (
		retv *ModuleInfo
		sts = []string{coreAddModePath, user, pass}
	)
	err := cli.Send(sts,&retv)
	if err != nil {
		return nil, err
	}
	return retv,nil
}