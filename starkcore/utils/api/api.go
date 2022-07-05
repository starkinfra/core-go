package api

import (
	"core-go/starkcore/utils/case"
	"core-go/starkcore/utils/resource"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetJson(response *http.Response) struct{} {
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	var data struct{}
	json.Unmarshal(resBody, &data)
	return data
}

func PostJson(object struct{}) string {
	u, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
	return string(u)
}

func Endpoint(resource resource.Resource) string {
	name := strings.Replace(resource.Name, "-log", "/log", 1000000)
	name = strings.Replace(resource.Name, "-attempt", "/attempt", 1000000)
	return _case.CamelToKebab(name)
}

func LastName(resource resource.Resource) string {
	name := strings.SplitN(resource.Name, "-", -1)
	return _case.CamelToKebab(strings.Join(name, " "))
}

func LastNamePlural(resource resource.Resource) string {
	base := LastName(resource)
	if strings.HasSuffix(base, "s") == true {
		return base
	}
	if strings.HasSuffix(base, "y") == true && strings.HasSuffix(base, "ey") == false {
		return fmt.Sprintf("%bs", base[:-1])
	}
	return fmt.Sprintf("%bs", base)
}
