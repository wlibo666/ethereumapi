package core

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestGetBlockCount(t *testing.T) {
	ct, err := GetBlockCount(rpcAddr)
	if err != nil {
		t.Fatalf("GetBlockCount failed,err:%s", err.Error())
	}
	t.Logf("GetBlockCount:%d", ct.Int64())
}

func TestGetBlockByNumber(t *testing.T) {
	blk, err := GetBlockByNumber(rpcAddr, big.NewInt(25))
	if err != nil {
		t.Fatalf("GetBlockByNumber failed,err:%s", err.Error())
	}
	t.Logf("block:%v", blk)
	cxt, err := json.Marshal(blk)
	if err == nil {
		t.Logf("block json:%s", string(cxt))
	}
}

func TestGetBlockByHash(t *testing.T) {
	blk, err := GetBlockByHash(rpcAddr, common.HexToHash("0x2bb1464de26db1afb55d3908d5e0067cd3a14c2a9477a7b61924569e6a07ed4c"))
	if err != nil {
		t.Fatalf("GetBlockByHash failed,err:%s", err.Error())
	}
	t.Logf("block:%v", blk)
}

func TestGetTranscationByHash(t *testing.T) {
	trs, pend, err := GetTranscationByHash(rpcAddr, common.HexToHash("0x083837f1fad1b6ac0f3d77fd033335cbd099bffc33ed29ae9d88aeabeeb7ab20"))
	if err != nil {
		t.Fatalf("GetTranscationByHash failed,err:%s", err.Error())
	}
	t.Logf("pendFlag:%v,trans:%v", pend, trs)
}

func TestGetTransactionCountByNumber(t *testing.T) {
	cnt, err := GetTransactionCountByNumber(rpcAddr, big.NewInt(24))
	if err != nil {
		t.Fatalf("GetTransactionCountByNumber failed,err:%s", err.Error())
	}
	t.Logf("block index:24,transaction count:%d", cnt)
}

func TestGetTransactionCountByHash(t *testing.T) {
	cnt, err := GetTransactionCountByHash(rpcAddr, common.HexToHash("0x2fea5e28bbcfecc7450141c05c32615e938d54175f53a97f8ec86a1ff073604d"))
	if err != nil {
		t.Fatalf("GetTransactionCountByHash failed,err:%s", err.Error())
	}
	t.Logf("block transaction count:%d", cnt)
}

func TestGetPendingTransactionCount(t *testing.T) {
	cnt, err := GetPendingTransactionCount(rpcAddr)
	if err != nil {
		t.Fatalf("GetPendingTransactionCount failed,err:%s", err.Error())
	}
	t.Fatalf("GetPendingTransactionCount:%d", cnt)
}
