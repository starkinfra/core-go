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
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Status  int
	Content []byte
}

func Fetch(host string, sdkVersion string, user user.User, method string, path string, payload interface{}, apiVersion string, language string, timeout int, query map[string]interface{}) (Response, Error.StarkErrors) {
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
	agent := fmt.Sprintf("Go-SDK-%v-%v", host, sdkVersion)
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

	request, _ := client.Do(req)

	if request.StatusCode == 400 {
		err := Error.InputError(request.Body)
		return Response{}, err
	}
	if request.StatusCode == 500 {
		err := Error.InternalServerError()
		return Response{}, err
	}
	if request.StatusCode != 200 {
		err := Error.UnknownError(request.Body)
		return Response{}, err
	}
	resp, _ := ioutil.ReadAll(request.Body)
	response := Response{Status: request.StatusCode, Content: resp}
	return response, Error.StarkErrors{}
}
