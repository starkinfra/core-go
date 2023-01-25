package rest

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"github.com/starkinfra/core-go/starkcore/utils/request"
	"io/ioutil"
	"math"
	"strconv"
)

func GetPage(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, query map[string]interface{}) ([]byte, string, Error.StarkErrors) {
	data := map[string]interface{}{}
	page, err := request.Fetch(
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
	resp, _ := ioutil.ReadAll(page.Body)
	unmarshalError := json.Unmarshal(resp, &data)
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

func GetStream(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, query map[string]interface{}, c chan map[string]interface{}, e chan Error.StarkError) {
	var response []map[string]interface{}
	datas := make([]map[string]interface{}, len(response))
	var isNilCursor, isValid bool
	limitQuery := make(map[string]interface{})
	limit, _ := strconv.Atoi(fmt.Sprintf("%v", query["limit"]))
	if limit == 0 {
		isValid = true
		limitQuery["limit"] = nil
	}
	for k, v := range query {
		limitQuery[k] = v
	}
	go func() {
		defer close(c)
		if limit != 0 {
			limitQuery["limit"] = int(math.Min(float64(limit), 100))
		}
		for _ = 0; (limit > 0 && isNilCursor == false) || isValid == true && limitQuery["limit"] == nil; {
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
			if len(err.Errors) != 0 {
				e <- err.Errors[0]
				break
			}
			copy(datas, response)
			json.Unmarshal(entities, &datas)
			for _, data := range datas {
				c <- data
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
}

func GetId(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	get, err := request.Fetch(
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
	resp, _ := ioutil.ReadAll(get.Body)
	unmarshalError := json.Unmarshal(resp, &data)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(resource)])
	return jsonBytes, err
}

func GetContent(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string, subResourceName string, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	get, err := request.Fetch(
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
	if err.Errors != nil {
		return nil, err
	}
	content, _ := ioutil.ReadAll(get.Body)
	return content, err
}

func GetSubResource(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string, subResourceName map[string]string, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	get, err := request.Fetch(
		host,
		sdkVersion,
		user,
		"GET",
		fmt.Sprintf("%v/%v/%v",
			api.Endpoint(resource),
			id,
			api.Endpoint(subResourceName)),
		"",
		apiVersion,
		language,
		timeout,
		query,
	)
	if err.Errors != nil {
		return nil, err
	}
	resp, _ := ioutil.ReadAll(get.Body)
	unmarshalError := json.Unmarshal(resp, &data)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(subResourceName)])
	return jsonBytes, err
}

func PostMulti(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, entity interface{}, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	post, err := request.Fetch(
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
	resp, _ := ioutil.ReadAll(post.Body)
	return api.FromApiJson(resp, resource), err
}

func PostSingle(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, entity interface{}, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	post, err := request.Fetch(
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
	resp, _ := ioutil.ReadAll(post.Body)
	unmarshalError := json.Unmarshal(resp, &data)
	if unmarshalError != nil {
		return nil, Error.StarkErrors{}
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(resource)])
	return jsonBytes, err
}

func PostSubResource(sdkVersion string, host string, apiVersion string, user user.User, resource map[string]string, id string, subResource map[string]string, entity interface{}, language string, timeout int) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	post, err := request.Fetch(
		host,
		sdkVersion,
		user,
		"POST",
		fmt.Sprintf("%v/%v/%v",
			api.Endpoint(resource),
			id,
			subResource),
		api.ApiJson(entity, resource),
		apiVersion,
		language,
		timeout,
		nil,
	)
	if err.Errors != nil {
		return nil, err
	}
	resp, _ := ioutil.ReadAll(post.Body)
	unmarshalError := json.Unmarshal(resp, &data)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(resource)])
	return jsonBytes, err
}

func DeleteId(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	cancel, err := request.Fetch(
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
	resp, _ := ioutil.ReadAll(cancel.Body)
	unmarshalError := json.Unmarshal(resp, &data)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(resource)])
	return jsonBytes, err
}

func PatchId(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string, payload interface{}, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	data := map[string]interface{}{}
	update, err := request.Fetch(
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
	resp, _ := ioutil.ReadAll(update.Body)
	unmarshalError := json.Unmarshal(resp, &data)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(resource)])
	return jsonBytes, err
}

func GetRaw(sdkVersion string, host string, apiVersion string, language string, timeout int, path string, user user.User, query map[string]interface{}) (map[string]interface{}, Error.StarkErrors) {
	data := map[string]interface{}{}
	raw, err := request.Fetch(
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
	resp, _ := ioutil.ReadAll(raw.Body)
	unmarshalError := json.Unmarshal(resp, &data)
	if unmarshalError != nil {
		return nil, err
	}
	return data, err
}
