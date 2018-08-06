package core

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/json-iterator/go"
	"github.com/wlibo666/ethereumapi/contract"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type TokenInfo struct {
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Decimals    int64   `json:"decimal"`
	TotalSupply big.Int `json:"totalSupply"`
}

func (info TokenInfo) String() string {
	content, err := json.Marshal(info)
	if err != nil {
		return err.Error()
	}
	return string(content)
}

func GetTokenInfo(rpcAddr, contractAddr string) (TokenInfo, error) {
	tokenInfo := TokenInfo{}

	dial, err := rpc.Dial(rpcAddr)
	if err != nil {
		return tokenInfo, err
	}
	defer dial.Close()
	cli := ethclient.NewClient(dial)
	defer cli.Close()

	tk, err := contract.NewToken(common.HexToAddress(contractAddr), cli)
	if err != nil {
		return tokenInfo, err
	}

	tokenInfo.Name, err = tk.Name(nil)
	if err != nil {
		return tokenInfo, err
	}
	tokenInfo.Symbol, err = tk.Symbol(nil)
	if err != nil {
		return tokenInfo, err
	}
	decimals, err := tk.Decimals(nil)
	if err != nil {
		return tokenInfo, err
	}
	tokenInfo.Decimals = decimals.Int64()
	total, err := tk.TotalSupply(nil)
	if err != nil {
		return tokenInfo, err
	}
	tokenInfo.TotalSupply = *total
	return tokenInfo, nil
}

func GetToken(rpcAddr, contractAddr, accountAddr string) (big.Int, error) {
	amount := big.NewInt(0)
	dial, err := rpc.Dial(rpcAddr)
	if err != nil {
		return *amount, err
	}
	defer dial.Close()
	cli := ethclient.NewClient(dial)
	defer cli.Close()

	tk, err := contract.NewToken(common.HexToAddress(contractAddr), cli)
	if err != nil {
		return *amount, err
	}
	amount, err = tk.BalanceOf(nil, common.HexToAddress(accountAddr))
	return *amount, nil
}

func SendToken(rpcAddr, contractAddr string, jsonKey []byte, passphrase, toAddr string, amounts *big.Int) (string, error) {
	dial, err := rpc.Dial(rpcAddr)
	if err != nil {
		return "", err
	}
	defer dial.Close()
	cli := ethclient.NewClient(dial)
	defer cli.Close()

	tk, err := contract.NewToken(common.HexToAddress(contractAddr), cli)
	if err != nil {
		return "", err
	}
	trOpts, err := bind.NewTransactor(bytes.NewReader(jsonKey), passphrase)
	if err != nil {
		return "", err
	}
	tr, err := tk.TokenTransactor.Transfer(trOpts, common.HexToAddress(toAddr), amounts)
	if err != nil {
		return "", err
	}
	return tr.Hash().String(), nil
}
