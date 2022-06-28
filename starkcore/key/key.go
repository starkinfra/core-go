package key

import (
	"errors"
	"github.com/starkbank/ecdsa-go/ellipticcurve/curve"
	"github.com/starkbank/ecdsa-go/ellipticcurve/privatekey"
	"io/ioutil"
	"log"
	"os"
)
func Create(path string){

	private := privatekey.New(curve.Secp256k1)
	public := private.PublicKey()

	privatePem := private.ToPem()
	dataPrivate := []byte(privatePem)
	publicPem := public.ToPem()
	dataPublic := []byte(publicPem)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		_, e := os.Create("sample/private-key.pem")
		errPrivate := ioutil.WriteFile("sample/private-key.pem", dataPrivate, 0)

		if e != nil || errPrivate != nil {
			log.Println(e)
		}

		_, f := os.Create("sample/public-key.pem")
		errPublic := ioutil.WriteFile("sample/public-key.pem", dataPublic, 0)

		if f != nil || errPublic != nil {
			log.Println(e)
		}
	}

}