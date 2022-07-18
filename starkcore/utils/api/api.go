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
	var bodyDefinitivo = map[string]interface{}{}
	bodyDefinitivo["boletos"] = CastJsonToApiFormat(m)

	jsons, err := json.Marshal(bodyDefinitivo)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Sprintln(string(jsons))

	return string(jsons)
}

func CastJsonToApiFormat(m interface{}) interface{} {
	var val = reflect.ValueOf(m)
	var body = make(map[string]interface{})
	var bodySlice []map[string]interface{}

	if val.Kind() == reflect.Map {
		for _, e := range val.MapKeys() {
			v := val.MapIndex(e)
			myMap := make(map[interface{}]interface{})
			switch t := v.Interface().(type) {
			case []interface{}:
				for key, value := range t {
					myMap[key] = value
					return CastJsonToApiFormat(myMap)
				}
			case map[string]interface{}:
				for key, value := range t {
					key = strcase.ToLowerCamel(key)
					body[key] = value
					if value == nil {
						delete(body, key)
					}
				}
				bodySlice = append(bodySlice, body)
			}
		}
		return bodySlice
	}
	return nil
}

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

	return base
}
