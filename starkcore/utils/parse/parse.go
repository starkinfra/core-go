package parse

import (
	u "core-go/starkcore/user/user"
	"core-go/starkcore/utils/rest"
	"github.com/starkbank/ecdsa-go/ellipticcurve/publickey"
)

func ParseAndVerify() {
	return
}

func IsSignatureValid() bool {

	return false
}

func GetPublicKey(sdkVersion, host, apiVersion string, user u.Users, language string, timeout int, refresh bool) string {
	publicKey := ""
	if publicKey != "" && refresh == false {
		return publicKey
	}

	pem := rest.GetRaw(sdkVersion, host, apiVersion, user, language, timeout)
	publicKey = publickey.FromPem(pem)
	cache["stark-public-key"] = publicKey
	return publicKey
}
