package request

import (
	"core-go/starkcore/environment"
	u "core-go/starkcore/user/user"
	"core-go/starkcore/utils/checks"
	"encoding/json"
	"fmt"
	"github.com/starkbank/ecdsa-go/ellipticcurve/ecdsa"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Fetch(host string, sdkVersion string, user u.Users, method string, path string, payload string, apiVersion string, language string) *http.Response {

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
	var agent = fmt.Sprintf("Golang-1.18-SDK-%v-%v", host, sdkVersion)

	if payload == "" {
		fmt.Println("BODY VAZIO")
	}

	var accessTime = strconv.FormatInt(time.Now().Unix(), 10)

	var message = fmt.Sprintf("%v:%v:%v", "project/", accessTime, payload)

	var signature = ecdsa.Sign(message, u.PrivateKey(user)).ToBase64()

	fmt.Printf("\nMESSAGE: %v\n", message)

	client := http.Client{Timeout: time.Duration(15) * time.Second}

	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		log.Fatal(err)
		fmt.Println("deu erro aqui")
	}

	req.Header.Add("Access-Id", "project/")
	req.Header.Add("Access-Time", accessTime)
	req.Header.Add("Access-Signature", signature)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", agent)
	req.Header.Add("Accept-Language", language)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}

	var m = map[string]interface{}{}
	if err := json.Unmarshal(resBody, &m); err != nil {
		panic(err)
	}
	//fmt.Println(sec)
	//
	//var m = make(map[string]string)
	//json.Unmarshal(resBody, &m)

	fmt.Println(m)
	//fmt.Println(errsdc)

	//var data struct{}
	//json.Unmarshal(resBody, &data)

	return res
}
