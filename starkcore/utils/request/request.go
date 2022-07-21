package request

import (
	"core-go/starkcore/environment"
	errors "core-go/starkcore/error"
	u "core-go/starkcore/user/user"
	"core-go/starkcore/utils/checks"
	"fmt"
	"github.com/starkbank/ecdsa-go/ellipticcurve/ecdsa"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Fetch(host string, sdkVersion string, user u.Users, method string, path string, payload string, apiVersion string, language string, timeout int) *http.Response {

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

	//https://development.api.starkinfra.com/v2/static-brcode

	//agent := fmt.Sprintf("Golang-1.%m-SDK-%h-%s", goversion.Version, host, sdkVersion)
	var agent = fmt.Sprintf("Golang-SDK-%v-%v", host, sdkVersion)

	var accessTime = strconv.FormatInt(time.Now().Unix(), 10)

	var message = fmt.Sprintf("%v:%v:%v", "project/", accessTime, payload)

	var signature = ecdsa.Sign(message, u.PrivateKey(user)).ToBase64()

	client := http.Client{Timeout: time.Duration(timeout) * time.Second}

	req, err := http.NewRequest(method, url, strings.NewReader(payload))

	req.Header.Add("Access-Id", "project/")
	req.Header.Add("Access-Time", accessTime)
	req.Header.Add("Access-Signature", signature)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", agent)
	req.Header.Add("Accept-Language", language)

	response, err := client.Do(req)
	if err != nil {
		switch response.StatusCode {
		case 400:
			errors.InputError(err)
		case 500:
			errors.InternalServerError()
		default:
			return response
		}
	}

	return response
}
