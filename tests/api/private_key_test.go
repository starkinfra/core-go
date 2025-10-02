package api

import (
	"github.com/starkinfra/core-go/starkcore/key"
	"testing"
)

func TestCreatePrivateKey(t *testing.T) {
	privateKey, publicKey := key.Create("")
	if privateKey == "" || publicKey == "" {
		t.Errorf("privateKey or publicKey is empty")
	}
}

func TestPathPrivateKey(t *testing.T) {
	privateKey, publicKey := key.Create("sample")
	if privateKey == "" || publicKey == "" {
		t.Errorf("privateKey or publicKey is empty")
	}
}
