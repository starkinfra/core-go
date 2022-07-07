package api

import (
	"bytes"
	"core-go/starkcore/utils/case"
	"core-go/starkcore/utils/subresource"
	"encoding/json"
	"fmt"
	"io"
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

func ToApi(payload io.Reader) *bytes.Reader {
	body, _ := json.Marshal(payload)
	qualquermerda := bytes.NewReader(body)
	fmt.Println(string(body))
	return qualquermerda
}

func Endpoint(resource subresource.Subresource) string {
	name := strings.Replace(resource.Name, "-log", "/log", 1000000)
	name = strings.Replace(resource.Name, "-attempt", "/attempt", 1000000)
	return _case.CamelToKebab(name)
}

func LastName(resource subresource.Subresource) string {
	name := strings.SplitN(resource.Name, "-", -1)
	return _case.CamelToKebab(strings.Join(name, " "))
}

func LastNamePlural(resource subresource.Subresource) string {
	base := LastName(resource)
	if strings.HasSuffix(base, "s") == true {
		return base
	}
	if strings.HasSuffix(base, "y") == true && strings.HasSuffix(base, "ey") == false {
		return fmt.Sprintf("%bs", base[:1])
	}
	return fmt.Sprintf("%bs", base)
}
