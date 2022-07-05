package user

import "github.com/starkbank/ecdsa-go/ellipticcurve/privatekey"

type User interface {
	AccessId() string
}

type Users struct {
	Id          string
	Pem         string
	Environment string
}

func PrivateKey(user Users) *privatekey.PrivateKey {
	var result = privatekey.FromPem(user.Pem)
	return &result
}
