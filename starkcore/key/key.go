package key

import (
	"fmt"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/curve"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/privatekey"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Create(path string) (string, string) {

	//	Generate a new key pair
	//
	//	Generates a secp256k1 ECDSA private/public key pair to be used in the API authentications
	//
	//	Parameters (optional):
	//	- path [string]: path to save the keys .pem files. No files will be saved if this parameter isn't provided
	//
	//	Return:
	//	- private and public key pems

	private := privatekey.New(curve.Secp256k1)
	public := private.PublicKey()

	privatePem := private.ToPem()
	publicPem := public.ToPem()

	if path != "" {
		makeDir(path, privatePem, publicPem)
	}
	return privatePem, publicPem
}

func makeDir(path, privatePem, publicPem string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		mkdirError := os.MkdirAll(path, os.ModePerm)
		if mkdirError != nil {
			panic(mkdirError)
		}
	}
	_, createPrivateKeyError := os.Create(filepath.Join(path, "private-key.pem"))
	_, createPublicKeyError := os.Create(filepath.Join(path, "public-key.pem"))

	if createPrivateKeyError != nil || createPublicKeyError != nil {
		fmt.Println(createPrivateKeyError, createPublicKeyError)
	}

	privateKeyError := ioutil.WriteFile(filepath.Join(path, "private-key.pem"), []byte(privatePem), 0666)
	publicKeyError := ioutil.WriteFile(filepath.Join(path, "public-key.pem"), []byte(publicPem), 0666)
	if privateKeyError != nil || publicKeyError != nil {
		fmt.Println(privateKeyError, publicKeyError)
	}
}
