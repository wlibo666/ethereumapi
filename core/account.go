package core

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

const (
	KEYSTORE_DIR = "/tmp"
)

func NewAccount(passphrase string) ([]byte, error) {
	ks := keystore.NewKeyStore(KEYSTORE_DIR, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(passphrase)
	if err != nil {
		return []byte{}, err
	}
	return ks.Export(account, passphrase, passphrase)
}

func UpdateAccount(keyJSON []byte, passphrase, newPassphrase string) ([]byte, error) {
	ks := keystore.NewKeyStore(KEYSTORE_DIR, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.Import(keyJSON, passphrase, passphrase)
	if err != nil {
		return []byte{}, err
	}
	err = ks.Update(account, passphrase, newPassphrase)
	if err != nil {
		return []byte{}, err
	}
	return ks.Export(account, newPassphrase, newPassphrase)
}
