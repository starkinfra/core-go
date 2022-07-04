package request

import (
	"core-go/starkcore/environment"
	"core-go/starkcore/user/project"
	"core-go/starkcore/utils/checks"
	"encoding/json"
	"fmt"
	"github.com/starkbank/ecdsa-go/ellipticcurve/ecdsa"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Status  string
	Content string
}

func GetJson(response *http.Response) struct{} {
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	var data struct{}
	json.Unmarshal(resBody, &data)
	return data
}

func Fetch(host string, sdkVersion string, user project.Projects, method string, path string, payload io.Reader, apiVersion string, query any, language string) *http.Response {

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

	//agent := fmt.Sprintf("Golang-1.%m-SDK-%h-%s", goversion.Version, host, sdkVersion)
	agent := fmt.Sprintf("Golang-1.0-SDK-%h-%s", host, sdkVersion)

	accessTime := string(time.Now().Unix())

	body := payload
	ERROBIZARRO message := fmt.Sprintf("%a:%t:%b", "", accessTime, body)
	ERROBIZARRO signature := ecdsa.Sign(message, "")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	req.Header = http.Header{
		"Access-Id":        {ERROBIZARRO},
		"Access-Time":      {accessTime},
		"Access-Signature": {ERROBIZARRO},
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
