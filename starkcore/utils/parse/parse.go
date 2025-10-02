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
	Error "github.com/starkinfra/core-go/starkcore/error"
)

func ParseAndVerify(content string, signature string, sdkVersion string, apiVersion string, language string, timeout int, host string, user user.User, key string) (interface{}, Error.StarkErrors) {
	jsonParse := map[string]interface{}{}
	verifiedContent, err := Verify(content, signature, sdkVersion, apiVersion, language, timeout, host, user)
	if err.Errors != nil {
		return nil, err
	}
	if key != "" {
		for _, value := range jsonParse {
			jsonParse[key] = value
		}
	}
	return verifiedContent, Error.StarkErrors{}
}

func Verify(content string, signature string, sdkVersion string, apiVersion string, language string, timeout int, host string, user user.User) (interface{}, Error.StarkErrors) {
	var public publickey.PublicKey
	signatureFromBase64 := Signature.FromBase64(signature)
	if signatureFromBase64.ToBase64() == "" {
		return nil, error.InvalidSignatureError("The provided signature is not valid")
	}

	publicKey, err := getPublicKey(
		sdkVersion,
		host,
		apiVersion,
		language,
		timeout,
		user,
		false,
	)
	if err.Errors != nil {
		return nil, err
	}
	public = publicKey.(publickey.PublicKey)

	if isSignatureValid(content, signatureFromBase64, public) {
		return content, Error.StarkErrors{}
	}

	publicKey, err = getPublicKey(
		sdkVersion,
		host,
		apiVersion,
		language,
		timeout,
		user,
		true,
	)
	public = publicKey.(publickey.PublicKey)

	if isSignatureValid(content, signatureFromBase64, public) {
		return content, Error.StarkErrors{}
	}

	return nil, error.InvalidSignatureError("The provided signature and content do not match the public key")
}

func isSignatureValid(content string, signature Signature.Signature, publicKey publickey.PublicKey) bool {
	if ecdsa.Verify(content, signature, &publicKey) {
		return true
	}
	return false
}

func getPublicKey(sdkVersion, host, apiVersion string, language string, timeout int, user user.User, refresh bool) (interface{}, Error.StarkErrors) {
	mapPem := make(map[string]interface{})
	data := map[string]interface{}{}

	publicKey := cache.Cache["stark-public-key"]
	if publicKey != nil && refresh == false {
		return publicKey, Error.StarkErrors{}
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
		"",
		true,
	)
	if err.Errors != nil {
		return nil, err
	}

	unmarshalError := json.Unmarshal(pem.Content, &data)
	if unmarshalError != nil {
		return nil, Error.InputError(string(pem.Content))
	}

	json.Unmarshal([]byte(fmt.Sprintf("%v", data["cursor"])), &mapPem)
	content := fmt.Sprintf("%v", data["publicKeys"].([]interface{})[0].(map[string]interface{})["content"])
	publicKey = publickey.FromPem(content)
	cache.Cache["stark-public-key"] = publicKey
	return publicKey, Error.StarkErrors{}
}
