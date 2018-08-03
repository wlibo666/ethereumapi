package core

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetETCByAddr(rpcAddr, AccountAddr string) (big.Int, error) {
	c, err := rpc.Dial(rpcAddr)
	if err != nil {
		return *big.NewInt(0), err
	}
	defer c.Close()

	client := ethclient.NewClient(c)
	defer client.Close()

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(AccountAddr), nil)
	if err != nil {
		return *big.NewInt(0), err
	}
	return *balance, nil
}

func SendETC(rpcAddr string, jsonKey []byte, passphrase string, toAddr string, amount *big.Int,
	gasLimit uint64, gasPrice *big.Int, data []byte) (string, error) {
	key, err := keystore.DecryptKey(jsonKey, passphrase)
	if err != nil {
		return "", err
	}
	c, err := rpc.Dial(rpcAddr)
	if err != nil {
		return "", err
	}
	defer c.Close()

	client := ethclient.NewClient(c)
	defer client.Close()

	nonce, err := client.NonceAt(context.Background(), key.Address, nil)
	if err != nil {
		return "", err
	}
	tmpTrans := types.NewTransaction(nonce, common.HexToAddress(toAddr), amount, gasLimit, gasPrice, data)
	trans, err := types.SignTx(tmpTrans, types.HomesteadSigner{}, key.PrivateKey)
	err = client.SendTransaction(context.Background(), trans)
	if err != nil {
		return trans.Hash().String(), err
	}
	return trans.Hash().String(), nil
}
