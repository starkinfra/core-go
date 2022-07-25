package user

import "github.com/starkbank/ecdsa-go/ellipticcurve/privatekey"

type User interface {
	AccessId() string
	Environments() string
	PrivateKeys() *privatekey.PrivateKey
}

type Users struct {
	Id          string
	Pem         string
	Environment string
}
