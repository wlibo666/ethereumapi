package core

import (
	"testing"
)

var (
	mypwd    = "mypassphrase"
	mynewpwd = "mynewpassphrase"
)

func TestNewAccount(t *testing.T) {
	jsonKey, err := NewAccount(mypwd)
	if err != nil {
		t.Fatalf("NewAccount failed,err:%s", err.Error())
	}
	t.Logf("account jsonKey:%s\n\n", string(jsonKey))

	newJsonKey, err := UpdateAccount(jsonKey, mypwd, mynewpwd)
	if err != nil {
		t.Fatalf("UpdateAccount failed,err:%s", err.Error())
	}
	t.Logf("account new jsonKey:%s\n\n", string(newJsonKey))
}
