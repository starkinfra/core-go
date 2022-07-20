package api

import (
	"core-go/starkcore/utils/case"
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"net/http"
	"reflect"
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

func ApiJson(payload interface{}, resource map[string]string) string {
	var m = map[string]interface{}{}

	out, _ := json.Marshal(payload)
	json.Unmarshal(out, &m)

	jsons, err := json.Marshal(CastJsonToApiFormat(m))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return string(jsons)
}

func CastJsonToApiFormat(m interface{}) interface{} {
	var mJson = reflect.ValueOf(m)
	if mJson.Kind() != reflect.Map {
		return m
	}

	apiJson := map[string]interface{}{}

	if typedJson, ok := mJson.Interface().(map[string]interface{}); ok {
		for key, value := range typedJson {
			if value == nil {
				continue
			}
			key = strcase.ToLowerCamel(key)

			if v, ok := value.([]interface{}); ok {
				jsonSlice := []interface{}{}
				for _, val := range v {
					jsonSlice = append(jsonSlice, CastJsonToApiFormat(val))
				}
				apiJson[key] = jsonSlice
				continue
			}
			if v, ok := value.(map[string]interface{}); ok {
				apiJson[key] = CastJsonToApiFormat(v)
				continue
			}
			apiJson[key] = value
		}
	}
	return apiJson
}

func Endpoint(resource map[string]string) string {
	name := strings.Replace(resource["name"], "-log", "/log", 1000000)
	name = strings.Replace(resource["name"], "-attempt", "/attempt", 1000000)
	return name
}

func LastName(resource string) string {
	name := strings.SplitN(resource, "-", -1)
	return _case.CamelToKebab(strings.Join(name, " "))
}

func LastNamePlural(resource string) string {
	base := LastName(resource)
	if strings.HasSuffix(base, "s") == true {
		return base
	}
	if strings.HasSuffix(base, "s") == false {
		ok := base + "s"
		return ok
	}
	if strings.HasSuffix(base, "y") == true && strings.HasSuffix(base, "ey") == false {
		return fmt.Sprintf("%bs", base[:1])
	}

	return base
}
