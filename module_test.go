package gomsf

import "testing"

func TestModAuxiliarys(t *testing.T) {
	_, err := ModAuxiliarys(Cli, Token)
	if err != nil {
		t.Error(err)
		return
	}
	//t.Logf("auxs: %s",rets)
}

func TestModExploits(t *testing.T) {
	_, err := ModExploits(Cli, Token)
	if err != nil {
		t.Error(err)
		return
	}
	//t.Logf("exploits: %s",rets)
}

func TestModPayloads(t *testing.T) {
	_, err := ModPayloads(Cli, Token)
	if err != nil {
		t.Error(err)
		return
	}
	//t.Logf("payloads: %s",rets)
}

func TestModNops(t *testing.T) {
	rets, err := ModNops(Cli, Token)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("nops: %s", rets)
}
func TestModPosts(t *testing.T) {
	_, err := ModPosts(Cli, Token)
	if err != nil {
		t.Error(err)
		return
	}
	//t.Logf("posts: %s",rets)
}

func TestModEncoderss(t *testing.T) {
	rets, err := ModEncoders(Cli, Token)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("encoders: %s", rets)
}

func TestModOptions(t *testing.T) {
	rets, err := ModExploits(Cli, Token)
	if err != nil {
		t.Error(err)
		return
	}
	options, err := OptionsOfMod(Cli, Token, ExploitType, rets.Module[0])
	if err != nil {
		t.Error(err)
		return
	}
	for i := range options {
		t.Logf("module:%s, options %v", rets.Module[0], options[i])
	}
}

func TestModinfo(t *testing.T) {
	rets, err := ModExploits(Cli, Token)
	if err != nil {
		t.Error(err)
		return
	}
	info, err := InfoOfMod(Cli, Token, ExploitType, rets.Module[0])
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("module:%s, info %v", rets.Module[0], info)
}
