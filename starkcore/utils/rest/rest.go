package rest

import (
	"core-go/starkcore/user/user"
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/request"
	"fmt"
	//"math"
	"net/http"
)

func GetPage(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, language string, timeout int) (interface{}, interface{}) {
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
	)
	response := api.FromApi(json)
	cursor := response
	return response, cursor
}

//func GetStream(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, language string, timeout int) map[string]interface{} {
//	if limit == nil {
//		var limitQuery map[string]interface{}
//	}
//	limitQuery = math.Min(limit, 100)
//	var json = request.Fetch(
//		host,
//		sdkVersion,
//		user,
//		"GET",
//		api.Endpoint(resource),
//		"",
//		apiVersion,
//		language,
//		timeout,
//	)
//	var response = api.FromApi(json)
//	return response
//}

func GetId(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, id string, language string, timeout int) string {
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
	)
	response := api.FromApi(json)
	return response
}

func GetContent(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, id string, language string, subResourceName string, timeout int) *http.Response {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v/%v", api.Endpoint(resource), id, subResourceName),
		"",
		apiVersion,
		language,
		timeout,
	)
	return json
}

func GetSubResource(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, id string, language string, subResourceName string, timeout int) string {
	json := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v/%v", api.Endpoint(resource), id, subResourceName),
		"",
		apiVersion,
		language,
		timeout,
	)
	response := api.FromApi(json)
	return response
}

func PostMulti(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, payload string, language string, timeout int) interface{} {
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
	)
	response := api.FromApi(json)
	return response
}

func PostSingle(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, payload string, language string, timeout int) interface{} {
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
	)
	response := api.FromApi(json)
	return response
}

func DeleteId(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, id string, payload string, language string, timeout int) interface{} {
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
	)
	response := api.FromApi(json)
	return response
}

func PatchId(sdkVersion string, host string, apiVersion string, user user.Users, resource map[string]string, id string, payload string, language string, timeout int) interface{} {
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
	)
	response := api.FromApi(json)
	return response
}

func GetRaw(sdkVersion string, host string, apiVersion string, user user.Users, language string, timeout int) string {
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
	)
	return api.FromApi(json)
}
