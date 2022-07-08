package request

import (
	"core-go/starkcore/environment"
	u "core-go/starkcore/user/user"
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/checks"
	"fmt"
	"github.com/starkbank/ecdsa-go/ellipticcurve/ecdsa"
	"io"
	"log"
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

	var baseUrl = environment.Environments{
		Production: fmt.Sprintf("https://api.%s.com/%v", service, apiVersion),
		Sandbox:    fmt.Sprintf("https://sandbox.api.%s.com/%v", service, apiVersion),
	}

	var url = ""
	if user.Environment == "production" {
		url = fmt.Sprintf("%v/%v", baseUrl.Production, path)
	}
	if user.Environment == "sandbox" {
		url = fmt.Sprintf("%v/%v", baseUrl.Sandbox, path)
	}

	// https://development.api.starkinfra.com/v2/static-brcode
	//var url = fmt.Sprintf("%v/%v", baseUrl, path)

	//agent := fmt.Sprintf("Golang-1.%m-SDK-%h-%s", goversion.Version, host, sdkVersion)
	var agent = fmt.Sprintf("Golang-1.0-SDK-%v-%v", host, sdkVersion)

	var accessTime string = string(time.Now().Unix())

	if payload != nil {
		api.ToApi(payload)
	} else {
		payload = nil
	}

	var message = fmt.Sprintf("%a:%t:%b", "project/XXXXXXXXX", accessTime, payload)

	var signature = ecdsa.Sign(message, u.PrivateKey(user)).ToBase64()

	fmt.Println(url)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		log.Fatalln(err)
		os.Exit(1)
	}

	req.Header = http.Header{
		"Access-Id":        {"project/XXXXXXXXXXX"},
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
	defer res.Body.Close()
	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	fmt.Println(req.Header)

	return res
}
