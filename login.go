package gomsf

const (
	authlogin = "auth.login"
	authlogout = "auth.logout"
	authTokenlist = "auth.token_list"
	authTokenadd  = "auth.token_add"
	authTokengen  = "auth.token_generate"
	authTokenrm   = "auth.token_remove"
)

type AuthLogin struct {
	Result string `msgpack:"result"`
	Token  string `msgpack:"token"`
}

type AuthList struct {
	Tokens []string `msgpack:"tokens"`
}

func Login(cli MsfCli, user, pass string) (string, error) {
	var (
		retv *AuthLogin
		sts = []string{authlogin, user, pass}
	)
	err := cli.Send(sts,&retv)
	if err != nil {
		return "", err
	}
	return retv.Token,nil
}

func Logout(cli MsfCli, token, oldtoken string ) error {
	var (
		retv *AuthLogin
		sts = []string{authlogout, token, oldtoken}
	)
	return cli.Send(sts,&retv)
}


func (m *MsfGo) TokenAdd(cli MsfCli, token, newtoken string) error {
	var (
		retv *AuthLogin
		sts = []string{authTokenadd, token, newtoken}
	)
	return cli.Send(sts,&retv)
}

func (m *MsfGo) TokenGen(cli MsfCli, token string) error {
	var (
		retv *AuthLogin
		sts = []string{authTokengen, token}
	)
	return cli.Send(sts,&retv)
}

func (m *MsfGo) TokenRm(cli MsfCli, token, rmtoken string) error {
	var (
		retv *AuthLogin
		sts = []string{authTokenrm, token, rmtoken}
	)
	return cli.Send(sts,&retv)
}

func (m *MsfGo) Tokenlist(cli MsfCli, token string) (*AuthList, error ){
	var (
		retv *AuthList
		sts = []string{authTokenlist, token}
	)
	return retv, cli.Send(sts,&retv)
}
