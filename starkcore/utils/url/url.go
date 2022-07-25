package url

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"reflect"
	"strings"
)

func UrlEncode(params map[string]interface{}) string {
	if params != nil {
		params = castQueryToApiFormat(params)
		for k, v := range params {
			params[k] = valueToString(v)
		}
		b, _ := json.Marshal(params)
		return fmt.Sprint("?", string(b))
	} else {
		return ""
	}
	return ""
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
				jsonSlice := []interface{}{}
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
