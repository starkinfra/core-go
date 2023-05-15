package rest

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"github.com/starkinfra/core-go/starkcore/utils/request"
	"math"
	"strconv"
)

func GetPage(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, query map[string]interface{}) ([]byte, string, Error.StarkErrors) {
	data := map[string]interface{}{}
	response, err := request.Fetch(
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
	if err.Errors != nil {
		return nil, "", err
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	cursor := data["cursor"]
	jsonBytes, _ := json.Marshal(data[api.LastNamePlural(resource)])
	if cursor == nil {
		return jsonBytes, "", err
	}
	return jsonBytes, cursor.(string), err
}

func GetStream(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, query map[string]interface{}) chan map[string]interface{} {
	channel := make(chan map[string]interface{})
	isNilCursor := false
	var response []map[string]interface{}
	newResponse := make([]map[string]interface{}, len(response))

	limitQuery := make(map[string]interface{})
	for k, v := range query {
		limitQuery[k] = v
	}
	limit, _ := strconv.Atoi(fmt.Sprintf("%v", query["limit"]))
	if limit == 0 {
		limitQuery["limit"] = nil
	}
	if limit != 0 {
		limitQuery["limit"] = int(math.Min(float64(limit), 100))
	}

	go func() {
		defer close(channel)
		for _ = 0; (limitQuery["limit"] == nil || limit > 0) && !isNilCursor; {
			entities, cursor, err := GetPage(
				sdkVersion,
				host,
				apiVersion,
				language,
				timeout,
				user,
				resource,
				limitQuery,
			)
			if err.Errors != nil {
				for _, e := range err.Errors {
					panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
				}
			}
			copy(newResponse, response)
			unmarshalErr := json.Unmarshal(entities, &newResponse)
			if unmarshalErr != nil {
				panic(unmarshalErr)
			}
			for _, data := range newResponse {
				channel <- data
			}
			if limit != 0 {
				limit -= 100
				limitQuery["limit"] = int(math.Min(float64(limit), 100))
			}
			limitQuery["cursor"] = cursor
			if cursor == "" {
				isNilCursor = true
			}
		}
	}()
	return channel
}

func GetId(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	response, err := request.Fetch(
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
	if err.Errors != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(resource)])
	return jsonBytes, err
}

func GetContent(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string, subResourceName string, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	response, err := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf(
			"%v/%v/%v",
			api.Endpoint(resource),
			id,
			subResourceName,
		),
		"",
		apiVersion,
		language,
		timeout,
		query,
	)
	if err.Errors != nil {
		return nil, err
	}
	content, _ := json.Marshal(response.Content)
	return content, err
}

func GetSubResource(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string, subResourceName map[string]string, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	response, err := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf(
			"%v/%v/%v",
			api.Endpoint(resource),
			id,
			api.Endpoint(subResourceName),
		),
		"",
		apiVersion,
		language,
		timeout,
		query,
	)
	if err.Errors != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(subResourceName)])
	return jsonBytes, err
}

func PostMulti(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, entity interface{}, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	response, err := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resource),
		api.ApiJson(entity, resource),
		apiVersion,
		language,
		timeout,
		query,
	)
	if err.Errors != nil {
		return nil, err
	}
	return api.FromApiJson(response.Content, resource), err
}

func PostSingle(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, entity interface{}, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	response, err := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		api.Endpoint(resource),
		api.ApiJson(entity, resource),
		apiVersion,
		language,
		timeout,
		query,
	)
	if err.Errors != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(resource)])
	return jsonBytes, err
}

func PostSubResource(sdkVersion string, host string, apiVersion string, user user.User, resource map[string]string, id string, subResource map[string]string, entity interface{}, language string, timeout int) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	response, err := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		fmt.Sprintf(
			"%v/%v/%v",
			api.Endpoint(resource),
			id,
			api.Endpoint(subResource),
		),
		api.ApiJson(entity, resource),
		apiVersion,
		language,
		timeout,
		nil,
	)
	if err.Errors != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(subResource)])
	return jsonBytes, err
}

func DeleteId(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	response, err := request.Fetch(
		host,
		sdkVersion,
		user,
		"DELETE",
		fmt.Sprintf("%v/%v", api.Endpoint(resource), id),
		"",
		apiVersion,
		language,
		timeout,
		query,
	)
	if err.Errors != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	resp, _ := json.Marshal(data[api.LastName(resource)])
	return resp, err
}

func PatchId(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string, payload interface{}, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	response, err := request.Fetch(
		host,
		sdkVersion,
		user,
		"PATCH",
		fmt.Sprintf("%v/%v", api.Endpoint(resource), id),
		api.ApiJson(payload, resource),
		apiVersion,
		language,
		timeout,
		query,
	)
	if err.Errors != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(resource)])
	return jsonBytes, err
}

func GetRaw(sdkVersion string, host string, apiVersion string, language string, timeout int, path string, user user.User, query map[string]interface{}) (map[string]interface{}, Error.StarkErrors) {
	data := map[string]interface{}{}
	response, err := request.Fetch(
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
	if err.Errors != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	return data, err
}

func PostRaw(sdkVersion string, host string, apiVersion string, language string, timeout int, path string, payload interface{}, user user.User, query map[string]interface{}) (map[string]interface{}, Error.StarkErrors) {
	data := map[string]interface{}{}
	response, err := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		path,
		payload,
		apiVersion,
		language,
		timeout,
		query,
	)
	if err.Errors != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	return data, err
}
