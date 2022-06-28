package api

import (
	"fmt"
	"strings"
	"core-go/starkcore/utils/case"
)

func ApiJson(entity string) {
	type := fmt.Sprintf("%T", entity)
	if ()
}

func CastJsonToApiFormat(){
	return _case.SnakeToCamel()
}

func CastValues(value string){

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

