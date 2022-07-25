package parse

import (
	error2 "core-go/starkcore/error"
	u "core-go/starkcore/user/user"
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/cache"
	"core-go/starkcore/utils/rest"
	json2 "encoding/json"
	"fmt"
	"github.com/starkbank/ecdsa-go/ellipticcurve/ecdsa"
	"github.com/starkbank/ecdsa-go/ellipticcurve/publickey"
	signature2 "github.com/starkbank/ecdsa-go/ellipticcurve/signature"
)

func ParseAndVerify(content string, signature signature2.Signature, sdkVersion string, apiVersion string, host string, resource string, user u.User, language string, timeout int, key string) string {
	content = Verify(content, signature, sdkVersion, apiVersion, host, language, user, timeout)
	jsonParse := map[string]interface{}{}
	json2.Unmarshal([]byte(content), &jsonParse)

	if key != "" {
		for _, v := range jsonParse {
			jsonParse[key] = v
		}
	}

	return api.FromApi(resource, jsonParse)
}

func Verify(content string, signature signature2.Signature, sdkVersion string, apiVersion string, language string, host string, user u.User, timeout int) string {

	signature = signature2.FromBase64(signature.ToBase64())
	if signature != nil {
		error2.InvalidSignatureError("The provide signature is not valid")
	}

	publicKey := getPublicKey(
		sdkVersion,
		host,
		apiVersion,
		user,
		language,
		timeout,
		false,
	)

	if isSignatureValid(content, signature, publicKey) {
		return content
	}

	publicKey = getPublicKey(
		sdkVersion,
		host,
		apiVersion,
		user,
		language,
		timeout,
		true,
	)

	if isSignatureValid(content, signature, publicKey) {
		return content
	}

	return error2.InvalidSignatureError("The provided signature and content do not match the public key")
}

func isSignatureValid(content string, signature signature2.Signature, publicKey publickey.PublicKey) bool {

	if ecdsa.Verify(content, signature, &publicKey) {
		return true
	}

	normalized := map[string]string{}
	err := json2.Unmarshal([]byte(content), &normalized)
	// FALTA O .SORT
	if err != nil {
		return false
	}

	if ecdsa.Verify(normalized, signature, &publicKey) {
		return true
	}
	return false
}

func getPublicKey(sdkVersion, host, apiVersion string, user u.User, language string, timeout int, refresh bool) interface{} {
	mapPem := map[string]interface{}{}
	publicKey := cache.Cache["starkcore-public-key"]
	if publicKey != nil && refresh == false {
		return publicKey
	}

	pem := rest.GetRaw(
		sdkVersion,
		host,
		"/public-key",
		apiVersion,
		user,
		language,
		timeout,
		nil,
	)

	json2.Unmarshal([]byte(pem), &mapPem)
	sPem := fmt.Sprintf("%v", mapPem["publicKeys"])

	publicKey = publickey.FromPem(sPem)
	cache.Cache["stark-public-key"] = publicKey

	return publicKey
}
