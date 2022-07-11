package rest

import (
	"core-go/starkcore/user/user"
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/request"
	"fmt"
	"reflect"
)

func GetPage(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, language string) (struct{}, struct{}) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		"",
		apiVersion,
		language,
	)
	var response = api.FromApi(json)
	r := reflect.ValueOf(response)
	typeOfS := r.Type()

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("%s\t%v\n", typeOfS.Field(i).Name, r.Field(i).Interface())
	}

	fmt.Println(response)

	cursor := response
	return cursor, cursor

}

func GetStream(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, language string) struct{} {
	//if limit == nil {
	//	var limitQuery = any
	//}
	//var limitQuery = math.Min(limit, 100)
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		"",
		apiVersion,
		language,
	)
	var response = api.FromApi(json)
	return response
	//entityJson := response[api.LastName(resource)]
	//return api.FromApiJson(resource, entityJson), nil
}

func GetId(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, id string, language string) struct{} {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v", api.Endpoint(resource), id),
		"",
		apiVersion,
		language,
	)
	var response = api.FromApi(json)
	return response
	//entityJson := response[api.LastName(resource)]
	//return api.FromApiJson(resource, entityJson), nil
}

func GetContent(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, id string, language string, subResourceName string) struct{} {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v/%v", api.Endpoint(resource), id, subResourceName),
		"",
		apiVersion,
		language,
	)
	var response = api.FromApi(json)
	return response
}

func GetSubResource(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, language string) struct{} {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		"",
		apiVersion,
		language,
	)
	var response = api.FromApi(json)
	return response
	//entityJson := response[api.LastName(resource)]
	//return api.ToApi(entityJson), nil
}

func PostMulti(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, payload string, language string) (struct{}, struct{}) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resource),
		payload,
		apiVersion,
		language,
	)
	var response = api.FromApi(json)
	return response, struct{}{}
	//entityJson := response[api.LastName(resource)]
	//return api.ToApi(entityJson), nil
}

func PostSingle(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, payload string, language string) (struct{}, struct{}) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resource),
		payload,
		apiVersion,
		language,
	)

	var response = api.FromApi(json)
	r := reflect.ValueOf(response)
	typeOfS := r.Type()

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("%s\t%v\n", typeOfS.Field(i).Name, r.Field(i).Interface())
	}

	cursor := response

	return cursor, cursor
	//var response = api.FromApi(json)
	//return response, curos
	//entityJson := response[api.LastName(resource)]
	//return api.ToApi(entityJson), nil
}

func DeleteId(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, id string, payload string, language string) struct{} {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"DELETE",
		fmt.Sprintf("%e/%i", api.Endpoint(resource), id),
		payload,
		apiVersion,
		language,
	)
	var response = api.FromApi(json)
	return response
	//return api.ToApi(entityJson), nil
}

func PatchId(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, id string, payload string, language string) struct{} {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"PATCH",
		fmt.Sprintf("%e/%i", api.Endpoint(resource), id),
		payload,
		apiVersion,
		language,
	)
	var response = api.FromApi(json)
	return response
	//entity := response[api.LastName(resource)]
	//return api.ToApi(entityJson), nil
}

func GetRaw(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, language string) struct{} {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		"",
		apiVersion,
		language,
	)
	return api.FromApi(json)
}
