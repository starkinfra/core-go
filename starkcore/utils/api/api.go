package api

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"core-go/starkcore/utils/case"
	"time"
)

func ApiJson(entity io.Reader) map[]{
	if (reflect.TypeOf(entity).Name() == 'byte'){
		return CastJsonToApiFormat(entity)
	}

	json := map[]
	return CastJsonToApiFormat(json)
}

func CastJsonToApiFormat(json io.Reader) map[] {
	return _case.SnakeToCamel()
}

func CastValues(value io.Reader) map[string]string {
	accessTime := string(time.Now().Unix())
	switch value {
	case time.UnixDate:
	}

	var castedValues map[string]string


	return castedValues
}

func FromApiJson() string {

}

func Endpoint(resource map[string]string) string{
	name := resource["name"].ReplaceAllString("-log", "/log")
	name = resource["name"].ReplaceAllString("-attempt", "/attempt")
	return _case.CamelToKebab(name)
}

func LastName(resource map[string]string) string{
	name := resource["name"]
	nameS := strings.SplitN(name, "-", -1)
	return _case.CamelToKebab(nameS)
}

func LastNamePlural(resource map[string]string) string{
	base := LastName(resource)
	if (strings.HasSuffix(base, "s") == true){
		return base
	}
	if (strings.HasSuffix(base, "y") == true && strings.HasSuffix(base, "ey") == false){
		return fmt.Sprintf("%bs", base[:-1])
	}
	return fmt.Sprintf("%bs", base)
}

