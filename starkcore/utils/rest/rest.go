package rest

import (
	"core-go/starkcore/user/user"
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/request"
	"core-go/starkcore/utils/resource"
	"fmt"
	"reflect"
)

func GetPage(sdkVersion string, host string, apiVersion string, user user.Users, resource resource.Resource, language string) (any, any) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		nil,
		apiVersion,
		language,
	)
	var response = api.GetJson(json)
	r := reflect.ValueOf(response)
	typeOfS := r.Type()

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("%s\t%v\n", typeOfS.Field(i).Name, r.Field(i).Interface())
	}

	cursor := response

	return cursor, response
}

func GetStream(sdkVersion string, host string, apiVersion string, user user.Users, resource resource.Resource, language string) (any, any) {
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
		nil,
		apiVersion,
		language,
	)
	var response = api.GetJson(json)
	return response, nil
	//entityJson := response[api.LastName(resource)]
	//return api.FromApiJson(resource, entityJson), nil
}

func GetId(sdkVersion string, host string, apiVersion string, user user.Users, resource resource.Resource, id string, language string) (any, any) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v", api.Endpoint(resource), id),
		nil,
		apiVersion,
		language,
	)
	var response = api.GetJson(json)
	return response, nil
	//entityJson := response[api.LastName(resource)]
	//return api.FromApiJson(resource, entityJson), nil
}

func GetContent(sdkVersion string, host string, apiVersion string, user user.Users, resource resource.Resource, id string, language string, subResourceName string) (any, any) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v/%v", api.Endpoint(resource), id, subResourceName),
		nil,
		apiVersion,
		language,
	)
	var response = api.GetJson(json)
	return response, nil
}

func GetSubResource(sdkVersion string, host string, apiVersion string, user user.Users, resource resource.Resource, language string) (any, any) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		nil,
		apiVersion,
		language,
	)
	var response = api.GetJson(json)
	return response, nil
	//entityJson := response[api.LastName(resource)]
	//return api.PostJson(entityJson), nil
}

func GetSubResources(sdkVersion string, host string, apiVersion string, user user.Users, resource resource.Resource) (any, any) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		nil,
		apiVersion,
		"language",
	)
	var response = api.GetJson(json)
	return response, nil
	//entityJson := response[api.LastName(resource)]
	//return api.PostJson(entityJson), nil
}

func PostMulti(sdkVersion string, host string, apiVersion string, user user.Users, resource resource.Resource, payload struct{}, language string) (any, any) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resource),
		api.PostJson(payload),
		apiVersion,
		language,
	)
	var response = api.GetJson(json)
	return response, nil
	//entityJson := response[api.LastName(resource)]
	//return api.PostJson(entityJson), nil
}

func PostSingle(sdkVersion string, host string, apiVersion string, user user.Users, resource resource.Resource, payload struct{}, language string) (any, any) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resource),
		api.PostJson(payload),
		apiVersion,
		language,
	)
	var response = api.GetJson(json)
	return response, nil
	//entityJson := response[api.LastName(resource)]
	//return api.PostJson(entityJson), nil
}

func DeleteId(sdkVersion string, host string, apiVersion string, user user.Users, resource resource.Resource, id string, payload struct{}, language string) (any, any) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"DELETE",
		fmt.Sprintf("%e/%i", api.Endpoint(resource), id),
		api.PostJson(payload),
		apiVersion,
		language,
	)
	var response = api.GetJson(json)
	return response, nil
	//return api.PostJson(entityJson), nil
}

func PatchId(sdkVersion string, host string, apiVersion string, user user.Users, resource resource.Resource, id string, payload struct{}, language string) (any, any) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"PATCH",
		fmt.Sprintf("%e/%i", api.Endpoint(resource), id),
		api.PostJson(payload),
		apiVersion,
		language,
	)
	var response = api.GetJson(json)
	return response, nil
	//entity := response[api.LastName(resource)]
	//return api.PostJson(entityJson), nil
}

func GetRaw(sdkVersion string, host string, apiVersion string, user user.Users, resource resource.Resource, language string) (struct{}, any) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		nil,
		apiVersion,
		language,
	)
	return api.GetJson(json), nil
}
