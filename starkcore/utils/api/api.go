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
	var data = map[string]interface{}{}
	var resBody, _ = ioutil.ReadAll(response.Body)
	json.Unmarshal(resBody, &data)
	var rJson, _ = json.MarshalIndent(data, "", "  ")

	return string(rJson)
}

func ApiJson(payload interface{}) string {
	m := map[string]interface{}{}
	out, _ := json.Marshal(payload)
	json.Unmarshal(out, &m)
	jsons, _ := json.Marshal(CastJsonToApiFormat(m))

	return string(jsons)
}

func CastJsonToApiFormat(m interface{}) interface{} {
	var apiJson = map[string]interface{}{}
	var mJson = reflect.ValueOf(m)
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
				var jsonSlice = []interface{}{}
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
	var name = strings.Replace(resource["name"], "-log", "/log", 1000000)
	name = strings.Replace(resource["name"], "-attempt", "/attempt", 1000000)
	return name
}

func LastName(resource string) string {
	var name = strings.SplitN(resource, "-", -1)
	return strcase.ToKebab(strings.Join(name, " "))
}

func LastNamePlural(resource string) string {
	var base = LastName(resource)
	if strings.HasSuffix(base, "s") == true {
		return base
	}
	if strings.HasSuffix(base, "y") == true && strings.HasSuffix(base, "ey") == false {
		return base[:len(base)-1] + "ies"
	}
	return base + "s"
}
