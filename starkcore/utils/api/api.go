package api

import (
	"bytes"
	"core-go/starkcore/utils/case"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func FromApi(response *http.Response) struct{} {
	resBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}

	var data struct{}
	json.Unmarshal(resBody, &data)
	return data
}

func ToApi(payload interface{}) *bytes.Reader {
	fmt.Println("------------ENTRANDO NO ToApi-------------")
	body, _ := json.Marshal(payload)
	qualquermerda := bytes.NewReader(body)
	fmt.Sprintf("\nBODY PARSEADO: \n", string(body))
	return qualquermerda
}

//func ToApiString(payload io.Reader) string {
//	body, _ := json.Marshal(payload)
//	return string(body)
//}

func Endpoint(resource map[string]string) string {
	name := strings.Replace(resource["name"], "-log", "/log", 1000000)
	name = strings.Replace(resource["name"], "-attempt", "/attempt", 1000000)
	return name
}

func LastName(resource map[string]string) string {
	name := strings.SplitN(resource["name"], "-", -1)
	return _case.CamelToKebab(strings.Join(name, " "))
}

func LastNamePlural(resource map[string]string) string {
	base := LastName(resource)
	if strings.HasSuffix(base, "s") == true {
		return base
	}
	if strings.HasSuffix(base, "y") == true && strings.HasSuffix(base, "ey") == false {
		return fmt.Sprintf("%bs", base[:1])
	}
	return fmt.Sprintf("%bs", base)
}
