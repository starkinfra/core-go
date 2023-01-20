package user

import "github.com/starkbank/ecdsa-go/v2/ellipticcurve/privatekey"

type User interface {
	GetAcessId() string
	GetEnvironment() string
	GetPrivateKey() *privatekey.PrivateKey
}
