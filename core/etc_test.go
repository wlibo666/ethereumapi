package core

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	rpcAddr  = "http://xxx:8545"
	recvAddr = "0x399A918E65efae0D26E67D545129355Ccf090d3c"
	jsonKey  = `{"address":"49e7888acb220790b363e7061a8a9b46d58bfdc8","crypto":{"cipher":"aes-128-ctr","ciphertext":"34188705fedff14b83e0e069c95bcbc890fd09365173ec008b156f36c79829e5","cipherparams":{"iv":"40922e3a89d0d8daf938427f652b0995"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"205d1f722a0e555935d70a170135ecaa0eb6439eddb628196f6c2dc00d5eb431"},"mac":"274bfc314261b55181b6ef7e44b7578fcb8ea94f0491dd021e7516988618b916"},"id":"57328043-0e90-4298-bba6-5657731543e0","version":3}`
	pwd      = "www.xxx.com"
)

func TestGetETCByAddr(t *testing.T) {
	balance, err := GetETCByAddr(rpcAddr, recvAddr)
	if err != nil {
		t.Fatalf("GetETCByAddr failed,err:%s", err.Error())
	}
	t.Logf("GetETCByAddr for %s,etc:%s\n", recvAddr, balance.String())
}

func TestSendETC(t *testing.T) {
	amount, _ := hexutil.DecodeBig("0xDE0B6B3A7640000")
	hashAddr, err := SendETC(rpcAddr, []byte(jsonKey), pwd, recvAddr, amount, 1000000, big.NewInt(1), []byte{})
	if err != nil {
		t.Fatalf("sendETC failed,err:%s", err.Error())
	}
	t.Logf("trans addr:%s", hashAddr)
	balance, err := GetETCByAddr(rpcAddr, recvAddr)
	if err != nil {
		t.Fatalf("TestSendETC: GetETCByAddr failed,err:%s", err.Error())
	}
	t.Logf("TestSendETC: GetETCByAddr for %s,etc:%s\n", recvAddr, balance.String())
}
