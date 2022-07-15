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

//HOW TO PRETTY PRINT
//b, err := json.MarshalIndent(m, "", "  ")
//if err != nil {
//	fmt.Println("error:", err)
//}
//fmt.Print(string(b))

//fmt.Println(m)
//HOW TO RANGE A MAP
//for key, value := range m {
//if value == nil || value == "" {
//	fmt.Printf("%v", value)
//	delete(m, key)
//}
//	fmt.Printf("%v:%v", key, value)
//}

//fmt.Printf("JSON: %v", string(json))

func FromApi(response *http.Response) struct{} {
	resBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}

	var data struct{}
	json.Unmarshal(resBody, &data)
	return data
}

func ApiJson(payload interface{}) string {

	var m = map[string]interface{}{}

	out, _ := json.Marshal(payload)
	json.Unmarshal(out, &m)

	json, err := json.Marshal(CastJsonToApiFormat(m))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Sprintln(string(json))

	return string(json)
}

func CastJsonToApiFormat(m interface{}) interface{} {
	val := reflect.ValueOf(m)
	//bodyPost := map[interface{}]interface{}{}
	//bodySlice := []interface{}{}
	body := map[string]interface{}{}

	if val.Kind() == reflect.Map {
		for _, e := range val.MapKeys() {
			v := val.MapIndex(e)
			myMap := make(map[interface{}]interface{})
			switch t := v.Interface().(type) {
			case []interface{}:
				for key, value := range t {
					myMap[key] = value
					CastJsonToApiFormat(myMap)
				}
			case map[string]interface{}:
				for key, value := range t {
					key = strcase.ToLowerCamel(key)
					body[key] = value
					if value == nil {
						delete(body, key)
					}
				}
			}
		}
	}
	return body
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
	return fmt.Sprintf("%bs", base)
}
