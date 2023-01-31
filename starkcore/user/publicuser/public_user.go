package publicuser

import (
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/privatekey"
)

type PublicUser struct {
	Environment string
}

func (o PublicUser) GetAcessId() string {
	return ""
}

func (o PublicUser) GetEnvironment() string {
	return ""
}

func (o PublicUser) GetPrivateKey() *privatekey.PrivateKey {
	return nil
}
