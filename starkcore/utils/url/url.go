package url

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"net/url"
	"reflect"
	"strings"
)

func UrlEncode(params map[string]interface{}) string {
	if params == nil {
		return ""
	}
	p := url.Values{}
	for k, v := range params {
		if fmt.Sprintf("%v", reflect.TypeOf(v)) == "bool" {
			params[k] = fmt.Sprintf("%v", reflect.ValueOf(v))
		}
	}
	params = castQueryToApiFormat(params)
	for k, v := range params {
		p.Add(k, valueToString(v))
	}
	query := p.Encode()
	return fmt.Sprint("?" + query)
}

func valueToString(value interface{}) string {
	qJson := reflect.ValueOf(value)
	if sliceVal, ok := qJson.Interface().([]string); ok {
		result := fmt.Sprintf(strings.Join(sliceVal, ","))
		return result
	}
	if mapVal, ok := qJson.Interface().(map[string]interface{}); ok {
		b, _ := json.Marshal(mapVal)
		return string(b)
	}
	if stringVal, ok := qJson.Interface().(string); ok {
		return stringVal
	}
	if intVal, ok := qJson.Interface().(int); ok {
		return fmt.Sprintf("%v", intVal)
	}
	return ""
}

func castQueryToApiFormat(m interface{}) map[string]interface{} {
	apiJson := map[string]interface{}{}
	mJson := reflect.ValueOf(m)
	if typedJson, ok := mJson.Interface().(map[string]interface{}); ok {
		for key, value := range typedJson {
			if value == nil {
				continue
			}
			key = strcase.ToLowerCamel(key)
			if v, ok := value.([]interface{}); ok {
				var jsonSlice []interface{}
				for _, val := range v {
					jsonSlice = append(jsonSlice, castQueryToApiFormat(val))
				}
				apiJson[key] = jsonSlice
				continue
			}
			if v, ok := value.(map[string]interface{}); ok {
				apiJson[key] = castQueryToApiFormat(v)
				continue
			}
			apiJson[key] = value
		}
	}
	return apiJson
}
