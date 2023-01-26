package request

import (
	"encoding/json"
	"fmt"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/ecdsa"
	"github.com/starkinfra/core-go/starkcore/environment"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
	urls "github.com/starkinfra/core-go/starkcore/utils/url"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Fetch(host string, sdkVersion string, user user.User, method string, path string, payload interface{}, apiVersion string, language string, timeout int, query map[string]interface{}) (*http.Response, Error.StarkErrors) {
	var url string
	var body string
	language = checks.CheckLanguage(language)

	if payload != "" {
		bytes, _ := json.Marshal(payload)
		body = string(bytes)
	}

	switch user.GetEnvironment() {
	case environment.Environments.Production:
		url = fmt.Sprintf("https://api.stark%v.com/%v", host, apiVersion)
	case environment.Environments.Sandbox:
		url = fmt.Sprintf("https://sandbox.api.stark%v.com/%v", host, apiVersion)
	}

	url = fmt.Sprintf("%v/%v%v", url, path, urls.UrlEncode(query))
	agent := fmt.Sprintf("Golang-SDK-%v-%v", host, sdkVersion)
	accessTime := strconv.FormatInt(time.Now().Unix(), 10)
	message := fmt.Sprintf("%v:%v:%v", user.GetAcessId(), accessTime, body)
	signature := ecdsa.Sign(message, user.GetPrivateKey()).ToBase64()
	client := http.Client{Timeout: time.Duration(timeout) * time.Second}

	req, _ := http.NewRequest(method, url, strings.NewReader(body))

	req.Header.Add("Access-Id", user.GetAcessId())
	req.Header.Add("Access-Time", accessTime)
	req.Header.Add("Access-Signature", signature)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", agent)
	req.Header.Add("Accept-Language", language)

	response, _ := client.Do(req)

	if response.StatusCode == 400 {
		err := Error.InputError(response.Body)
		return nil, err
	}
	if response.StatusCode == 500 {
		err := Error.InternalServerError()
		return nil, err
	}
	if response.StatusCode != 200 {
		err := Error.UnknownError(response.Body)
		return nil, err
	}
	return response, Error.StarkErrors{}
}
