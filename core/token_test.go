package core

import (
	"math/big"
	"testing"
)

var (
	contractAddr = "0xcc33f3073f3e645f1a8ca094098cca68d8c1087c"
	ownerAddr    = "0x49e7888acb220790b363e7061a8a9b46d58bfdc8"
	tokenAddr    = "0xeFEc1Cc477AF130d7a34e3A9bd2AC546a00Fa2c6"
)

func TestGetTokenInfo(t *testing.T) {
	tinfo, err := GetTokenInfo(rpcAddr, contractAddr)
	if err != nil {
		t.Fatalf("GetTokenInfo failed,err:%s", err.Error())
	}
	t.Logf("tinfo:%s\n", tinfo)
}

func TestGetToken(t *testing.T) {
	amount, err := GetToken(rpcAddr, contractAddr, tokenAddr)
	if err != nil {
		t.Fatalf("GetToken failed,err:%s", err.Error())
	}
	t.Logf("amount for:%s is :%s", tokenAddr, amount.String())
}

func TestSendToken(t *testing.T) {
	amount, err := GetToken(rpcAddr, contractAddr, ownerAddr)
	if err != nil {
		t.Fatalf("before sendToken, GetToken failed,err:%s", err.Error())
	}
	t.Logf("before sendToken, amount for:%s is :%s", ownerAddr, amount.String())

	hashAddr, err := SendToken(rpcAddr, contractAddr, []byte(jsonKey), pwd, tokenAddr, big.NewInt(100))
	if err != nil {
		t.Fatalf("SendToken failed,err:%s", err.Error())
	}
	t.Logf("hashAddr:%s\n", hashAddr)

	amount, err = GetToken(rpcAddr, contractAddr, ownerAddr)
	if err != nil {
		t.Fatalf("end sendToken, GetToken failed,err:%s", err.Error())
	}
	t.Logf("end sendToken, amount for:%s is :%s", ownerAddr, amount.String())

}
