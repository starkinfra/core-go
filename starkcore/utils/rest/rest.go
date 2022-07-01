package rest

import (
	"core-go/starkcore/user/project"
	"core-go/starkcore/user/users"
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/request"
	"core-go/starkcore/utils/resource"
	"fmt"
	"io"
)

type Response struct {

}

func GetPage(sdkVersion string, host string, apiVersion string, user project.Projects, resources resource.Resource, language string, timeout string, query io.Reader) map[string]string {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		nil,
		apiVersion,
		nil,
		language,
	)
	json = request.GetJson(json)
	entities := api.FromApiJson(resources, entity)  api.LastNamePlural(resources)
	cursor := json["cursor"]

	return cursor, entities
}

func GetStream(sdkVersion string, host string, apiVersion string, user project.Projects, resources map[string]string, language string, limit io.Reader, query string) {
	limitQuery := ""
	limitQuery := ""
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		nil,
		apiVersion,
		nil,
		language,
	)
	json = request.GetJson(json)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func GetId(sdkVersion string, host string, apiVersion string, user project.Projects, resources map[string]string, id string, language string, timeout string, query string) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v", api.Endpoint(resources), id),
		nil,
		apiVersion,
		nil,
		language,
)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func GetContent(sdkVersion string, host string, apiVersion string, user project.Projects, resources map[string]string, language string, timeout string, query string) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		nil,
		apiVersion,
		nil,
		language,
)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func GetSubResource(sdkVersion string, host string, apiVersion string, user project.Projects, resources map[string]string, language string, timeout string, query string) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		nil,
		apiVersion,
		nil,
		language,
)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func GetSubResources(sdkVersion string, host string, apiVersion string, user project.Projects, resources map[string]string, language string, timeout string, query string) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		nil,
		apiVersion,
		nil,
		language,
	)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func PostMulti(sdkVersion string, host string, apiVersion string, user project.Projects, resources map[string]string, payload io.Reader, language string, timeout string, query string) {
	payload := api.ApiJson(entity)
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resources),
		payload,
		apiVersion,
		nil,
		language,
	)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func PostSingle(sdkVersion string, host string, apiVersion string, user project.Projects, resources map[string]string, payload io.Reader, language string, timeout string, query string) {
	payload := api.ApiJson(entity)
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resources),
		payload,
		apiVersion,
		nil,
		language,
	)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func DeleteId(sdkVersion string, host string, apiVersion string, user project.Projects, resources map[string]string, id string, language string, timeout string, query string) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"DELETE",
		fmt.Sprintf("%e/%i", api.Endpoint(resources), id),
		api.CastJsonToApiFormat(nil),
		apiVersion,
		nil,
		language,
	)
	entity := json[api.LastName(resources)]
	return api.FromApiJson(resources, entity)
}

func PatchId(sdkVersion string, host string, apiVersion string, user project.Projects, resources map[string]string, id string, payload string, language string, timeout string, query string) string {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"PATCH",
		fmt.Sprintf("%e/%i", api.Endpoint(resources), id),
		api.CastJsonToApiFormat(payload),
		apiVersion,
		nil,
		language,
	)
	json = request.GetJson(*json)
	entity := json[api.LastName(resources)]
	return api.FromApiJson(resources, entity)
}

func GetRaw(sdkVersion string, host string, apiVersion string, user project.Projects, resources map[string]string, language string, timeout string, query string) string {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		nil,
		apiVersion,
		nil,
		language,
		)
	return request.GetJson(*json)
}
