package core

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetBlockCount(rpcAddr string) (big.Int, error) {
	count := big.NewInt(0)
	dial, err := rpc.Dial(rpcAddr)
	if err != nil {
		return *count, err
	}
	defer dial.Close()
	cli := ethclient.NewClient(dial)
	defer cli.Close()

	block, err := cli.BlockByNumber(context.Background(), nil)
	if err != nil {
		return *count, err
	}
	return *block.Number(), nil
}

func GetBlockByNumber(rpcAddr string, number *big.Int) (*types.Block, error) {
	dial, err := rpc.Dial(rpcAddr)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	cli := ethclient.NewClient(dial)
	defer cli.Close()

	return cli.BlockByNumber(context.Background(), number)
}

func GetBlockByHash(rpcAddr string, blkHash common.Hash) (*types.Block, error) {
	dial, err := rpc.Dial(rpcAddr)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	cli := ethclient.NewClient(dial)
	defer cli.Close()

	return cli.BlockByHash(context.Background(), blkHash)
}

func GetTranscationByHash(rpcAddr string, hash common.Hash) (*types.Transaction, bool, error) {
	dial, err := rpc.Dial(rpcAddr)
	if err != nil {
		return nil, false, err
	}
	defer dial.Close()
	cli := ethclient.NewClient(dial)
	defer cli.Close()
	return cli.TransactionByHash(context.Background(), hash)
}

func GetTransactionCountByHash(rpcAddr string, blkHash common.Hash) (int, error) {
	dial, err := rpc.Dial(rpcAddr)
	if err != nil {
		return 0, err
	}
	defer dial.Close()
	cli := ethclient.NewClient(dial)
	defer cli.Close()

	blk, err := cli.BlockByHash(context.Background(), blkHash)
	if err != nil {
		return 0, err
	}
	return blk.Transactions().Len(), nil
}

func GetTransactionCountByNumber(rpcAddr string, number *big.Int) (int, error) {
	dial, err := rpc.Dial(rpcAddr)
	if err != nil {
		return 0, err
	}
	defer dial.Close()
	cli := ethclient.NewClient(dial)
	defer cli.Close()

	blk, err := cli.BlockByNumber(context.Background(), number)
	if err != nil {
		return 0, err
	}
	return blk.Transactions().Len(), nil
}

func GetPendingTransactionCount(rpcAddr string) (uint, error) {
	dial, err := rpc.Dial(rpcAddr)
	if err != nil {
		return 0, err
	}
	defer dial.Close()
	cli := ethclient.NewClient(dial)
	defer cli.Close()

	return cli.PendingTransactionCount(context.Background())
}
