package rest

import (
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/request"
	"fmt"
)

func GetPage(sdkVersion, host, apiVersion, user, resources, language, timeout, payload, query) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		payload,
		apiVersion,
		query,
	)
	json = request.GetJson(json, Response)
	entities := api.FromApiJson(resources, entity)  api.LastNamePlural(resources)
	cursor := json.get("cursor")

	return cursor, entities

}

func GetStream(sdkVersion, host, apiVersion, user, resources, id, language, timeout, payload, query) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		PAYLOAD DE MULTI,
		apiVersion,
		query,
	)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func GetId(sdkVersion, host, apiVersion, user, resources, id, language, timeout, payload, query) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		PAYLOAD DE MULTI,
		apiVersion,
		query,
)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func GetContent(sdkVersion, host, apiVersion, user, resources, id, language, timeout, payload, query) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		PAYLOAD DE MULTI,
		apiVersion,
		query,
)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func GetSubResource(sdkVersion, host, apiVersion, user, resources, id, language, timeout, payload, query) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		PAYLOAD DE MULTI,
		apiVersion,
		query,
)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func GetSubResources(sdkVersion, host, apiVersion, user, resources, id, language, timeout, payload, query) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		PAYLOAD DE MULTI,
		apiVersion,
		query,
	)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func PostMulti(sdkVersion, host, apiVersion, user, resources, id, language, timeout, payload, query) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resources),
		PAYLOAD DE MULTI,
		apiVersion,
		query,
	)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func PostSingle(sdkVersion, host, apiVersion, user, resources, id, language, timeout, payload, query) {
	payload := api.ApiJson(entity)
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resources),
		payload,
		apiVersion,
		query,
	)
	json = request.GetJson(json, Response)
	entityJson := json[api.LastName(resources)]
	return api.FromApiJson(resources, entityJson)
}

func DeleteId(sdkVersion, host, apiVersion, user, resources, id, language, timeout, payload, query) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"DELETE",
		fmt.Sprintf("%e/%i", api.Endpoint(resources), id),
		api.CastJsonToApiFormat(payload),
		apiVersion,
		query,
	)
	entity := json[api.LastName(resources)]
	return api.FromApiJson(resources, entity)
}

func PatchId(sdkVersion, host, apiVersion, user, resources, id,language, timeout, payload, query) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"PATCH",
		fmt.Sprintf("%e/%i", api.Endpoint(resources), id),
		api.CastJsonToApiFormat(payload),
		apiVersion,
		query,
	)
	entity := json[api.LastName(resources)]
	return api.FromApiJson(resources, entity)
}

func GetRaw(sdkVersion, host, apiVersion, user, resources, language, timeout, payload, query) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resources),
		payload,
		apiVersion,
		query,
		)
	return request.GetJson(json)
}
