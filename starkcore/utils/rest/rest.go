package rest

import (
	"core-go/starkcore/user/project"
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/request"
	"core-go/starkcore/utils/resource"
	"fmt"
	"io"
	"math"
	"reflect"
)

func GetPage(sdkVersion string, host string, apiVersion string, user project.Projects, resource resource.Resource, language string, timeout string, query io.Reader) (any, any) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		nil,
		apiVersion,
		nil,
		language,
	)
	var response = request.GetJson(json)
	r := reflect.ValueOf(response)
	typeOfS := r.Type()

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("%s\t%v\n", typeOfS.Field(i).Name, r.Field(i).Interface())
	}

	//entities := api.FromApiJson(resource, response)
	cursor := response

	return cursor, entities
}

func GetStream(sdkVersion string, host string, apiVersion string, user project.Projects, resource resource.Resource, language string, limit any) {
	if limit == nil {
		var limitQuery = any
	}
	var limitQuery = math.Min(limit, 100)
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		nil,
		apiVersion,
		limitQuery,
		language,
	)
	var response = request.GetJson(json)
	entityJson := response[api.LastName(resource)]
	return api.FromApiJson(resource, entityJson)
}

func GetId(sdkVersion string, host string, apiVersion string, user project.Projects, resource resource.Resource, id string, language string) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v", api.Endpoint(resource), id),
		nil,
		apiVersion,
		nil,
		language,
	)
	var response = request.GetJson(json)
	entityJson := response[api.LastName(resource)]
	return api.FromApiJson(resource, entityJson)
}

func GetContent(sdkVersion string, host string, apiVersion string, user project.Projects, resource resource.Resource, id string, language string, subResourceName string) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v/%v", api.Endpoint(resource), id, subResourceName),
		nil,
		apiVersion,
		nil,
		language,
	)
	var response = request.GetJson(json)
	return response
}

func GetSubResource(sdkVersion string, host string, apiVersion string, user project.Projects, resource resource.Resource, language string, timeout string, query string) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		nil,
		apiVersion,
		nil,
		language,
	)
	var response = request.GetJson(json)
	entityJson := response[api.LastName(resource)]
	return api.FromApiJson(resource, entityJson)
}

func GetSubResources(sdkVersion string, host string, apiVersion string, user project.Projects, resource resource.Resource, language string, timeout string, query string) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		nil,
		apiVersion,
		nil,
		language,
	)
	var response = request.GetJson(json)
	entityJson := response[api.LastName(resource)]
	return api.FromApiJson(resource, entityJson)
}

func PostMulti(sdkVersion string, host string, apiVersion string, user project.Projects, resource resource.Resource, payload io.Reader, language string, timeout string, query string) {
	payload := api.ApiJson(entity)
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resource),
		payload,
		apiVersion,
		nil,
		language,
	)
	var response = request.GetJson(json)
	entityJson := response[api.LastName(resource)]
	return api.FromApiJson(resource, entityJson)
}

func PostSingle(sdkVersion string, host string, apiVersion string, user project.Projects, resource resource.Resource, payload io.Reader, language string, timeout string, query string) {
	payload := api.ApiJson(entity)
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resource),
		payload,
		apiVersion,
		nil,
		language,
	)
	var response = request.GetJson(json)
	entityJson := response[api.LastName(resource)]
	return api.FromApiJson(resource, entityJson)
}

func DeleteId(sdkVersion string, host string, apiVersion string, user project.Projects, resource resource.Resource, id string, language string, timeout string, query string) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"DELETE",
		fmt.Sprintf("%e/%i", api.Endpoint(resource), id),
		api.CastJsonToApiFormat(nil),
		apiVersion,
		nil,
		language,
	)
	var response = json[api.LastName(resource)]
	return api.FromApiJson(resource, response)
}

func PatchId(sdkVersion string, host string, apiVersion string, user project.Projects, resource resource.Resource, id string, payload string, language string, timeout string, query string) string {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"PATCH",
		fmt.Sprintf("%e/%i", api.Endpoint(resource), id),
		api.CastJsonToApiFormat(payload),
		apiVersion,
		nil,
		language,
	)
	var response = request.GetJson(json)
	entity := response[api.LastName(resource)]
	return api.FromApiJson(resource, entity)
}

func GetRaw(sdkVersion string, host string, apiVersion string, user project.Projects, resource resource.Resource, language string) struct{} {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		nil,
		apiVersion,
		nil,
		language,
	)
	return request.GetJson(json)
}
