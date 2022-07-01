package request

import (
	"core-go/starkcore/environment"
	"core-go/starkcore/error"
	"core-go/starkcore/user/project"
	"core-go/starkcore/user/users"
	"core-go/starkcore/utils/checks"
	"github.com/starkbank/ecdsa-go/ellipticcurve/ecdsa"
	// "github.com/starkbank/ecdsa-go/ellipticcurve/privatekey"
	"internal/goversion"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	"fmt"
)

type Response struct {
	Status  string
	Content string
}

func GetJson(response *http.Request) string {
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	fmt.Printf(string(resBody))

	return response.Referer()
}

func Fetch(host string, sdkVersion string, user project.Projects, method string, path string, payload io.Reader, apiVersion string, query io.Reader, language string) *http.Request {

	sdkVersion = "v2"

	user = checks.CheckUser(user)

	language = "en-US"
	language = checks.CheckLanguage(language)

	service := checks.CheckHost(host)

	urlEnv := environment.Environments{
		Production: fmt.Sprintf("https://api.%s%v.com/", service, apiVersion),
		Sandbox:    fmt.Sprintf("https://sandbox.api.%s%v.com/", service, apiVersion),
	}

	url := fmt.Sprintf("%b/%p%q", urlEnv, path, query)

	agent := fmt.Sprintf("Golang-1.%m-SDK-%h-%s", goversion.Version, host, sdkVersion)

	accessTime := string(time.Now().Unix())

	body := payload
	ERROBIZARRO message := fmt.Sprintf("%a:%t:%b", "", accessTime, body)
	ERROBIZARRO signature := ecdsa.Sign(message, "")

	resp, err := http.NewRequest(method, url, nil)
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
		"Access-Id":        {ERROBIZARRO},
		"Access-Time":      {accessTime},
		"Access-Signature": {ERROBIZARRO},
		"Content-Type":     {"application/json"},
		"User-Agent":       {agent},
		"Accept-Language":  {language},
	}

	return resp
}
