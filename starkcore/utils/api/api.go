package api

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"reflect"
	"strings"
	"time"
)

func ApiJson(payload interface{}, resource map[string]string) string {
	b, _ := json.Marshal(payload)
	unmarshalSliceError := json.Unmarshal(b, &payload)
	if unmarshalSliceError != nil {
		fmt.Println(unmarshalSliceError)
	}
	data := map[string]interface{}{}
	if _, ok := payload.([]interface{}); ok {
		data[LastNamePlural(resource)] = payload
		return ApiJson(data, resource)
	}
	tApi := map[string]interface{}{}
	unmarshalError := json.Unmarshal(b, &tApi)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	apiJson, _ := json.Marshal(CastJsonToApiFormat(tApi))
	return string(apiJson)
}

func CastJsonToApiFormat(tApi interface{}) interface{} {
	apiJson := map[string]interface{}{}
	mJson := reflect.ValueOf(tApi)
	if mJson.Kind() != reflect.Map {
		return tApi
	}
	if typedJson, ok := mJson.Interface().(map[string]interface{}); ok {
		for key, value := range typedJson {
			if value == nil {
				continue
			}
			key = strcase.ToLowerCamel(key)
			if v, ok := value.([]interface{}); ok {
				var jsonSlice []interface{}
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
		for k, v := range apiJson {
			date, err := time.Parse(time.RFC3339, fmt.Sprintf("%v", v))
			if err == nil {
				converted := ConvertDateTime(date)
				apiJson[k] = converted
			}
		}
	}
	return apiJson
}

func FromApiJson(response interface{}, resource map[string]string) []byte {
	data := map[string]interface{}{}
	unmarshalError := json.Unmarshal(response.([]byte), &data)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	jsonStr, _ := json.Marshal(data[LastNamePlural(resource)])
	return jsonStr
}

func Endpoint(resource map[string]string) string {
	last := strings.SplitN(resource["name"], "-", -1)
	name := strcase.ToKebab(strings.Join(last, " "))
	endpoint := strings.Replace(name, "-log", "/log", 1)
	endpoint = strings.Replace(endpoint, "-attempt", "/attempt", 1)
	return endpoint
}

func LastName(resource map[string]string) string {
	last := strings.SplitN(resource["name"], "-", -1)
	name := strings.Split(strcase.ToKebab(strings.Join(last, " ")), "-")
	return name[len(name)-1]
}

func LastNamePlural(resource map[string]string) string {
	plural := LastName(resource)
	if strings.HasSuffix(plural, "s") == true {
		return plural
	}
	if strings.HasSuffix(plural, "y") == true && strings.HasSuffix(plural, "ey") == false {
		return plural[:len(plural)-1] + "ies"
	}
	if strings.HasSuffix(plural, "-log") == true {
		return plural[len(plural)-3:] + "s"
	}
	return plural + "s"
}

func ConvertDateTime(date time.Time) string {
	if date.Format("2006-01-02") == "0001-01-01" {
		return ""
	}
	if date.Hour() != 0 || date.Minute() != 0 || date.Second() != 0 || date.Nanosecond() != 0 {
		return date.Format("2006-01-02T15:04:05.000000+00:00")
	}
	return date.Format("2006-01-02")
}
