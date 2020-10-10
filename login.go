package gomsf

const (
	authlogin     = "auth.login"
	authlogout    = "auth.logout"
	authTokenlist = "auth.token_list"
	authTokenadd  = "auth.token_add"
	authTokengen  = "auth.token_generate"
	authTokenrm   = "auth.token_remove"
)

type authLogin struct {
	Result string `msgpack:"result"`
	Token  string `msgpack:"token"`
}

func Login(cli MsfCli, user, pass string) (string, error) {
	var (
		retv *authLogin
		sts  = []string{authlogin, user, pass}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return "", err
	}
	return retv.Token, nil
}

func Logout(cli MsfCli, token, oldtoken string) error {
	var (
		retv *Generic
		sts  = []string{authlogout, token, oldtoken}
	)
	return cli.Send(sts, &retv)
}

func TokenAdd(cli MsfCli, token, newtoken string) error {
	var (
		retv *Generic
		sts  = []string{authTokenadd, token, newtoken}
	)
	return cli.Send(sts, &retv)
}

func TokenGen(cli MsfCli, token string) (string, error) {
	var (
		retv *authLogin
		sts  = []string{authTokengen, token}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return "", err
	}
	return retv.Token, nil
}

func TokenRm(cli MsfCli, token, rmtoken string) error {
	var (
		retv *Generic
		sts  = []string{authTokenrm, token, rmtoken}
	)
	return cli.Send(sts, &retv)
}

func Tokenlist(cli MsfCli, token string) ([]string, error) {
	var (
		retv struct {
			Tokens []string `msgpack:"tokens"`
		}
		sts = []string{authTokenlist, token}
	)
	var r = &retv
	err := cli.Send(sts, &r)
	if err != nil {
		return nil, err
	}
	return retv.Tokens, nil
}
