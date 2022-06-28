package users

import (
	"core-go/starkcore/utils/checks"
	"core-go/starkcore/utils/resource"
	"github.com/starkbank/ecdsa-go/ellipticcurve/privatekey"
)

type User struct {
	Id          resource.Resource
	Pem         checks.CheckPrivateKey()
	Environment checks.CheckEnvironment()
}

func PrivateKey(pem string, user User) string {
	User{Pem: pem}
	privateKey := string(privatekey.FromPem(user))
	return privateKey
}
