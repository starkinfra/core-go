package key

import (
	"github.com/starkbank/ecdsa-go/ellipticcurve/curve"
	"github.com/starkbank/ecdsa-go/ellipticcurve/privatekey"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Create(path string) (string, string) {
	//	Generate a new key pair
	//	Generates a secp256k1 ECDSA private/public key pair to be used in the API authentications
	//
	//	Parameters (optional):
	//	path [string]: path to save the keys .pem files. No files will be saved if this parameter isn't provided
	//
	//	Return:
	//	private and public key pems

	private := privatekey.New(curve.Secp256k1)
	public := private.PublicKey()

	privatePem := private.ToPem()
	publicPem := public.ToPem()

	if path != "" {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			erro := os.MkdirAll(path, os.ModePerm)
			if erro != nil {
				log.Println(erro)
			}
		}

		_, err := os.Create(filepath.Join(path, "private-key.pem"))
		_, err2 := os.Create(filepath.Join(path, "public-key.pem"))

		if err != nil || err2 != nil {
			log.Println(err)
		}

		errPrivate := ioutil.WriteFile(filepath.Join(path, "private-key.pem"), []byte(publicPem), 0666)
		errPublic := ioutil.WriteFile(filepath.Join(path, "public-key.pem"), []byte(publicPem), 0666)

		if errPrivate != nil || errPublic != nil {
			log.Println(err)
		}
	}

	return privatePem, publicPem
}
