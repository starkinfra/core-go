package request

import (
	"core-go/starkcore/environment"
	"core-go/starkcore/error"
	"core-go/starkcore/user/organization"
	"core-go/starkcore/user/users"
	"core-go/starkcore/utils/checks"
	"core-go/starkcore/utils/hosts"
	"encoding/json"
	"fmt"
	"github.com/starkbank/ecdsa-go/ellipticcurve/ecdsa"
	"internal/goversion"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Status  string
	Content string
}

func GetJson(response string, target interface{}) error.Error {
	return json.NewDecoder(response).Decode(target)
}

func Fetch(host string, sdkVersion string, user string, method string, path string, payload string, apiVersion string, query string) []byte {

	user = checks.CheckUser(user)
	language := checks.CheckLanguage("en-Us")

	service := hosts.StarkHost{
		Infra: "starkinfra",
		Bank:  "starkbank",
	}

	urlType := environment.Environment{
		Sandbox:    fmt.Sprintf("https://sandbox.api.%s%v.com/", service, apiVersion),
		Production: fmt.Sprintf("https://api.%s%v.com/", service, apiVersion),
	}

	url := fmt.Sprintf("%b/%p%q", urlType, path, query)

	agent := fmt.Sprintf("Golang-1.%m-SDK-%h-%s", goversion.Version, host, sdkVersion)

	accessTime := string(time.Now().Unix())
	body := payload
	message := fmt.Sprintf("%a:%t:%b", organization.AccessId(), accessTime, body)
	signature := ecdsa.Sign(message, users.PrivateKey())

	resp, err := http.NewRequest(method, url, body)
	if err != nil {
		if http.StatusInternalServerError == 500 {
			error.InternalServerError()
		}
		if http.StatusBadRequest == 400 {
			error.InputErrors()
		}
		if http.StatusOK == 200 {
			error.UnkownError()
		}
	}

	resp.Header = http.Header{
		"Access-Id":        {},
		"Access-Time":      {accessTime},
		"Access-Signature": {signature},
		"Content-Type":     {"application/json"},
		"User-Agent":       {agent},
		"Accept-Language":  {language},
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
