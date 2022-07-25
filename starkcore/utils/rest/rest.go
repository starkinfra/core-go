package rest

import (
	"core-go/starkcore/user/user"
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/request"
	"fmt"
	"math"
	"net/http"
)

func GetPage(sdkVersion string, host string, apiVersion string, user user.User, resource interface{}, language string, timeout int, query map[string]interface{}) (interface{}, interface{}) {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		api.Endpoint(resource),
		"",
		apiVersion,
		language,
		timeout,
		query,
	)
	entities := api.FromApi(json)
	cursor := entities
	return entities, cursor
}

func GetStream(sdkVersion string, host string, apiVersion string, user user.User, resource map[string]string, language string, timeout int, limit int, query map[string]interface{}) map[string]interface{} {
	limitQuery := map[string]interface{}{}

	if limit != 0 {
		limitQuery = map[string]interface{}{"limit": int(math.Min(float64(limit), 100))}
	}

	for {
		json := request.Fetch(
			host,
			sdkVersion,
			user,
			"GET",
			api.Endpoint(resource),
			"",
			apiVersion,
			language,
			timeout,
			limitQuery,
		)

		response := api.FromApi(json)
		cursor := response

		if limit != 0 {
			limit -= 100
			query["limit"] = int(math.Min(float64(limit), 100))
		}

		if cursor == "" || (limit != 0 && limit <= 0) {
			break
		}
	}
	return nil
}

func GetId(sdkVersion string, host string, apiVersion string, user user.User, resource map[string]string, id string, language string, timeout int, query map[string]interface{}) string {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v", api.Endpoint(resource), id),
		"",
		apiVersion,
		language,
		timeout,
		query,
	)
	entity := api.FromApi(json)
	return entity
}

func GetContent(sdkVersion string, host string, apiVersion string, user user.User, resource map[string]string, id string, language string, subResourceName string, timeout int, query map[string]interface{}) *http.Response {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v/%v",
			api.Endpoint(resource),
			id,
			subResourceName),
		"",
		apiVersion,
		language,
		timeout,
		query,
	)
	return json
}

func GetSubResource(sdkVersion string, host string, apiVersion string, user user.User, resource map[string]string, id string, language string, subResourceName string, timeout int, query map[string]interface{}) string {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v/%v",
			api.Endpoint(resource),
			id,
			subResourceName),
		"",
		apiVersion,
		language,
		timeout,
		query,
	)
	response := api.FromApi(json)
	return response
}

func PostMulti(sdkVersion string, host string, apiVersion string, user user.User, resource map[string]string, payload string, language string, timeout int, query map[string]interface{}) interface{} {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resource),
		payload,
		apiVersion,
		language,
		timeout,
		query,
	)
	response := api.FromApi(json)
	return response
}

func PostSingle(sdkVersion string, host string, apiVersion string, user user.User, resource map[string]string, payload string, language string, timeout int, query map[string]interface{}) interface{} {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resource),
		payload,
		apiVersion,
		language,
		timeout,
		query,
	)
	entityJson := api.FromApi(json)
	return entityJson
}

func DeleteId(sdkVersion string, host string, apiVersion string, user user.User, resource map[string]string, id string, payload string, language string, timeout int, query map[string]interface{}) interface{} {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"DELETE",
		fmt.Sprint(api.Endpoint(resource), "/", id),
		payload,
		apiVersion,
		language,
		timeout,
		query,
	)
	entity := api.FromApi(json)
	return entity
}

func PatchId(sdkVersion string, host string, apiVersion string, user user.User, resource map[string]string, id string, payload string, language string, timeout int, query map[string]interface{}) interface{} {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"PATCH",
		fmt.Sprint(api.Endpoint(resource), "/", id),
		payload,
		apiVersion,
		language,
		timeout,
		query,
	)
	entity := api.FromApi(json)
	return entity
}

func GetRaw(sdkVersion string, host string, path string, apiVersion string, user user.User, language string, timeout int, query map[string]interface{}) string {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		path,
		"",
		apiVersion,
		language,
		timeout,
		query,
	)
	return api.FromApi(json)
}
