package gomsf

import "testing"

func TestLogout(t *testing.T) {
	newt := "foo"
	err := TokenAdd(Cli, Token, newt)
	if err != nil {
		t.Error(err)
		return
	}
	err = Logout(Cli, Token, newt)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("success logout by token: %s", newt)
}

func TestTokenlist(t *testing.T) {
	newt, err := TokenGen(Cli, Token)
	if err != nil {
		t.Error(err)
		return
	}
	err = TokenAdd(Cli, Token, newt)
	if err != nil {
		t.Error(err)
		return
	}

	tokens, err := Tokenlist(Cli, Token)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		for i := range tokens {
			if tokens[i] != Token {
				TokenRm(Cli, Token, tokens[i])
			}
		}
	}()

	for i := range tokens {
		t.Logf("token-%d: %s", i, tokens[i])
	}
}
