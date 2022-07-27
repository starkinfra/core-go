package test

import (
	"core-go/starkcore/key"
	"fmt"
	"testing"
)

func TestCreatePrivateKey(t *testing.T) {

	privateKey, publicKey := key.Create("")
	fmt.Println("PRIVATE KEY", privateKey)
	fmt.Println("PUBLIC KEY", publicKey)

}

func TestPathPrivateKey(t *testing.T) {

	privateKey, publicKey := key.Create("sample")
	fmt.Println("PRIVATE KEY", privateKey)
	fmt.Println("PUBLIC KEY", publicKey)

}
