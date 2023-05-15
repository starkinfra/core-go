package parse

import (
	"encoding/json"
	"fmt"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/ecdsa"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/publickey"
	Signature "github.com/starkbank/ecdsa-go/v2/ellipticcurve/signature"
	"github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/cache"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
)

func ParseAndVerify(content string, signature string, sdkVersion string, apiVersion string, language string, timeout int, host string, user user.User, key string) interface{} {
	jsonParse := map[string]interface{}{}
	verifiedContent := Verify(content, signature, sdkVersion, apiVersion, language, timeout, host, user)
	if key != "" {
		for _, value := range jsonParse {
			jsonParse[key] = value
		}
	}
	return verifiedContent
}

func Verify(content string, signature string, sdkVersion string, apiVersion string, language string, timeout int, host string, user user.User) interface{} {
	var public publickey.PublicKey
	signatureFromBase64 := Signature.FromBase64(signature)
	if signatureFromBase64.ToBase64() == "" {
		panic(fmt.Sprintf("%v", error.InvalidSignatureError("The provided signature is not valid")))
	}

	publicKey := getPublicKey(
		sdkVersion,
		host,
		apiVersion,
		language,
		timeout,
		user,
		false,
	)
	public = publicKey.(publickey.PublicKey)

	if isSignatureValid(content, signatureFromBase64, public) == true {
		return content
	}

	publicKey = getPublicKey(
		sdkVersion,
		host,
		apiVersion,
		language,
		timeout,
		user,
		true,
	)
	public = publicKey.(publickey.PublicKey)

	if isSignatureValid(content, signatureFromBase64, public) == true {
		return content
	}

	panic(fmt.Sprintf("%v", error.InvalidSignatureError("The provided signature and content do not match the public key")))
}

func isSignatureValid(content string, signature Signature.Signature, publicKey publickey.PublicKey) bool {
	if ecdsa.Verify(content, signature, &publicKey) {
		return true
	}
	return false
}

func getPublicKey(sdkVersion, host, apiVersion string, language string, timeout int, user user.User, refresh bool) interface{} {
	mapPem := make(map[string]interface{})
	publicKey := cache.Cache["stark-public-key"]
	if publicKey != nil && refresh == false {
		return publicKey
	}

	pem, err := rest.GetRaw(
		sdkVersion,
		host,
		apiVersion,
		language,
		timeout,
		"/public-key",
		user,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	json.Unmarshal([]byte(fmt.Sprintf("%v", pem["cursor"])), &mapPem)
	content := fmt.Sprintf("%v", pem["publicKeys"].([]interface{})[0].(map[string]interface{})["content"])
	publicKey = publickey.FromPem(content)
	cache.Cache["stark-public-key"] = publicKey
	return publicKey
}
