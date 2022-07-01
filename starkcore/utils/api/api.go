package api

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"core-go/starkcore/utils/case"
)

func ApiJson(entity map[string]string) map[string]string{
	if (reflect.TypeOf(entity).Name() == 'map'){
		return CastJsonToApiFormat(entity)
	}

	json := map[]
	return CastJsonToApiFormat(json)
}

func CastJsonToApiFormat(json io.Reader) map[string]string {
	return _case.SnakeToCamel()
}

func Endpoint(resource map[string]string) string{
	name := strings.Replace(resource["name"],"-log", "/log", 1000000)
	name = strings.Replace(resource["name"],"-attempt", "/attempt", 1000000)
	return _case.CamelToKebab(name)
}

func LastName(resource map[string]string) string{
	name := strings.SplitN(resource["name"], "-", -1)
	return _case.CamelToKebab(strings.Join(name," "))
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

