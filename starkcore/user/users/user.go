package users

import (
	"github.com/starkbank/ecdsa-go/ellipticcurve/privatekey"
)

type User interface {
	PrivateKey() privatekey.PrivateKey
}

type Users struct {
	Id          string
	Pem         string
	Environment string
}

func PrivateKey(user Users) privatekey.PrivateKey {
	return privatekey.FromPem(user.Pem)
}
