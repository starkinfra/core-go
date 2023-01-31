package publicuser

import (
	"fmt"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/privatekey"
)

type PublicUser struct {
	Environment string
}

func (o PublicUser) GetAcessId() string {
	return ""
}

func (p PublicUser) GetEnvironment() string {
	return fmt.Sprintf("%v", p.Environment)
}

func (o PublicUser) GetPrivateKey() *privatekey.PrivateKey {
	return nil
}
