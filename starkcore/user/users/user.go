package users

import "github.com/starkbank/ecdsa-go/ellipticcurve/privatekey"

type User interface {
	AccessId() string
}

type Users struct {
	Id          string
	Pem         string
	Environment string
}

//func AcessId(user user.Users) string {
//	return project.AccessIdP(users.User)
//
//	return project.AccessIdO()
//}

func PrivateKey(user Users) privatekey.PrivateKey {
	return privatekey.FromPem(user.Pem)
}
