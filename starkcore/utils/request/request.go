package request

import (
	"core-go/starkcore/environment"
	u "core-go/starkcore/user/user"
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/checks"
	"fmt"
	"github.com/starkbank/ecdsa-go/ellipticcurve/ecdsa"
	"io"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Status  string
	Content string
}

func Fetch(host string, sdkVersion string, user u.Users, method string, path string, payload io.Reader, apiVersion string, language string) *http.Response {

	sdkVersion = "v2"

	language = "en-US"
	language = checks.CheckLanguage(language)

	var service = checks.CheckHost(host)

	var urlEnv = environment.Environments{
		Production: fmt.Sprintf("https://api.%s%v.com/", service, apiVersion),
		Sandbox:    fmt.Sprintf("https://sandbox.api.%s%v.com/", service, apiVersion),
	}

	var url = fmt.Sprintf("%b/%p%q", urlEnv, path)

	//agent := fmt.Sprintf("Golang-1.%m-SDK-%h-%s", goversion.Version, host, sdkVersion)
	var agent = fmt.Sprintf("Golang-1.0-SDK-%h-%s", host, sdkVersion)

	var accessTime = string(time.Now().Unix())

	var access u.User

	if payload == nil {
		api.ToApi(payload)
	} else {
		payload = nil
	}

	var message = fmt.Sprintf("%a:%t:%b", access.AccessId(), accessTime, payload)
	var signature = ecdsa.Sign(message, u.PrivateKey(user)).ToBase64()

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	req.Header = http.Header{
		"Access-Id":        {access.AccessId()},
		"Access-Time":      {accessTime},
		"Access-Signature": {signature},
		"Content-Type":     {"application/json"},
		"User-Agent":       {agent},
		"Accept-Language":  {language},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	return res
}
