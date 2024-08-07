package request

import (
	"encoding/json"
	"fmt"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/ecdsa"
	"github.com/starkinfra/core-go/starkcore/environment"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
	Url "github.com/starkinfra/core-go/starkcore/utils/url"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Status  int
	Content []byte
}

func Fetch(host string, sdkVersion string, user user.User, method string, path string, payload interface{}, apiVersion string, language string, timeout int, query map[string]interface{}, prefix string, throwError bool) (Response, Error.StarkErrors) {
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

	url = fmt.Sprintf("%v/%v%v", url, path, Url.UrlEncode(query))
	client := http.Client{Timeout: time.Duration(timeout) * time.Second}
	agent := fmt.Sprintf("Go-SDK-%v-%v", host, sdkVersion)
	if prefix != "" {
		agent = fmt.Sprintf("%v-Go-SDK-%v-%v", prefix, host, sdkVersion)
	}
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", agent)
	req.Header.Add("Accept-Language", language)
	req.Header.Add("Content-Type", "application/json")
	req = authenticationHeaders(user, body, req)

	rawResponse, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	responseContent, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		panic(err)
	}

	response := Response{Status: rawResponse.StatusCode, Content: responseContent}

	if throwError != false {
		if response.Status == 400 {
			err := Error.InputError(string(response.Content))
			return Response{}, err
		}
		if response.Status == 500 {
			err := Error.InternalServerError()
			return Response{}, err
		}
		if response.Status != 200 {
			err := Error.UnknownError(string(response.Content))
			return Response{}, err
		}
	}
	return response, Error.StarkErrors{}
}

func authenticationHeaders(user user.User, body string, req *http.Request) *http.Request {
	if reflect.TypeOf(user).Name() == "PublicUser" {
		return req
	}
	accessTime := strconv.FormatInt(time.Now().Unix(), 10)
	message := fmt.Sprintf("%v:%v:%v", user.GetAcessId(), accessTime, body)
	signature := ecdsa.Sign(message, user.GetPrivateKey()).ToBase64()
	req.Header.Add("Access-Id", user.GetAcessId())
	req.Header.Add("Access-Time", accessTime)
	req.Header.Add("Access-Signature", signature)
	return req
}
