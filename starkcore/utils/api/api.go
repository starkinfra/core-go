package api

import (
	"encoding/json"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func FromApi(response *http.Response) string {
	data := map[string]interface{}{}
	resBody, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(resBody, &data)
	rJson, _ := json.MarshalIndent(data, "", "  ")
	return string(rJson)
}

func ApiJson(payload interface{}) string {
	m := map[string]interface{}{}
	out, _ := json.Marshal(payload)
	json.Unmarshal(out, &m)
	apiJson, _ := json.Marshal(CastJsonToApiFormat(m))
	return string(apiJson)
}

func CastJsonToApiFormat(m interface{}) interface{} {
	apiJson := map[string]interface{}{}
	mJson := reflect.ValueOf(m)
	if mJson.Kind() != reflect.Map {
		return m
	}

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
	endpoint := strings.Replace(resource["name"], "-log", "/log", 1000000)
	endpoint = strings.Replace(resource["name"], "-attempt", "/attempt", 1000000)
	return endpoint
}

func LastName(resource string) string {
	last := strings.SplitN(resource, "-", -1)
	return strcase.ToKebab(strings.Join(last, " "))
}

func LastNamePlural(resource string) string {
	plural := LastName(resource)
	if strings.HasSuffix(plural, "s") == true {
		return plural
	}
	if strings.HasSuffix(plural, "y") == true && strings.HasSuffix(plural, "ey") == false {
		return plural[:len(plural)-1] + "ies"
	}
	return plural + "s"
}
