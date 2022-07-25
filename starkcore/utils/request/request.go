package request

import (
	"core-go/starkcore/environment"
	errors "core-go/starkcore/error"
	u "core-go/starkcore/user/user"
	"core-go/starkcore/utils/checks"
	url2 "core-go/starkcore/utils/url"
	"fmt"
	"github.com/starkbank/ecdsa-go/ellipticcurve/ecdsa"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Fetch(host string, sdkVersion string, user u.User, method string, path string, payload string, apiVersion string, language string, timeout int, query map[string]interface{}) *http.Response {
	url := ""
	sdkVersion = "v2"
	language = "en-US"
	user = checks.CheckUser(user)
	language = checks.CheckLanguage(language)
	service := checks.CheckHost(host)
	baseUrl := environment.Environments{
		Production: fmt.Sprintf("https://api.%v.com/%v", service, apiVersion),
		Sandbox:    fmt.Sprintf("https://sandbox.api.%v.com/%v", service, apiVersion),
	}

	if user.Environments() == "production" {
		url = fmt.Sprintf("%v/%v%v", baseUrl.Production, path, url2.UrlEncode(query))
	}
	if user.Environments() == "sandbox" {
		url = fmt.Sprintf("%v/%v%v", baseUrl.Sandbox, path, url2.UrlEncode(query))
	}

	agent := fmt.Sprintf("Golang-SDK-%v-%v", host, sdkVersion)
	accessTime := strconv.FormatInt(time.Now().Unix(), 10)
	message := fmt.Sprintf("%v:%v:%v", user.AccessId(), accessTime, payload)
	signature := ecdsa.Sign(message, user.PrivateKeys()).ToBase64()
	client := http.Client{Timeout: time.Duration(timeout) * time.Second}

	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
	}

	req.Header.Add("Access-Id", user.AccessId())
	req.Header.Add("Access-Time", accessTime)
	req.Header.Add("Access-Signature", signature)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", agent)
	req.Header.Add("Accept-Language", language)

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("ERRO", err)
		switch response.StatusCode {
		case 400:
			errors.InputError(response.Body)
		case 500:
			errors.InternalServerError()
		default:
			errors.UnknownError()
		}
	}

	return response
}
