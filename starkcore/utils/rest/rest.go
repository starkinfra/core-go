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
	jsonStr, _ := json.Marshal(data[api.LastNamePlural(resource)])
	if cursor == nil {
		return jsonStr, "", err
	}
	return jsonStr, cursor.(string), err
}

func GetStream(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	var jsonStr []byte
	starkErrors := Error.StarkErrors{}
	var entities, response []map[string]interface{}
	var isNilCursor, isValid bool
	limitQuery := make(map[string]interface{})
	c := make(chan map[string]interface{})
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
		for _ = 0; (limit > 0 && isNilCursor == false) || isValid == true; {
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
			isValid = false
			unmarshalError := json.Unmarshal(entities, &response)
			if unmarshalError != nil {
				_ = fmt.Errorf("%v", unmarshalError)
			}
			limitQuery["cursor"] = cursor
			if cursor == "" {
				isNilCursor = true
			}
			if limit != 0 {
				limit -= 100
				limitQuery["limit"] = int(math.Min(float64(limit), 100))
			}
			for _, data := range response {
				c <- data
			}
			starkErrors = err
		}
	}()
	for entity := range c {
		entities = append(entities, entity)
	}
	jsonStr, _ = json.Marshal(entities)
	return jsonStr, starkErrors
}

func GetId(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string) ([]byte, Error.StarkErrors) {
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
		nil,
	)
	if err.Errors != nil {
		return nil, err
	}
	resp, _ := ioutil.ReadAll(get.Body)
	unmarshalError := json.Unmarshal(resp, &data)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	jsonStr, _ := json.Marshal(data[api.LastName(resource)])
	return jsonStr, err
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
	jsonStr, _ := json.Marshal(data[api.LastName(subResourceName)])
	return jsonStr, err
}

func PostMulti(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, entity interface{}) ([]byte, Error.StarkErrors) {
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
		nil,
	)
	if err.Errors != nil {
		return nil, err
	}
	resp, _ := ioutil.ReadAll(post.Body)
	return api.FromApiJson(resp, resource), err
}

func PostSingle(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, entity interface{}) ([]byte, Error.StarkErrors) {
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
		nil,
	)
	if err.Errors != nil {
		return nil, err
	}
	resp, _ := ioutil.ReadAll(post.Body)
	unmarshalError := json.Unmarshal(resp, &data)
	if unmarshalError != nil {
		return nil, Error.StarkErrors{}
	}
	jsonStr, _ := json.Marshal(data[api.LastName(resource)])
	return jsonStr, err
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
	jsonStr, _ := json.Marshal(data[api.LastName(resource)])
	return jsonStr, err
}

func DeleteId(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string) ([]byte, Error.StarkErrors) {
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
		nil,
	)
	if err.Errors != nil {
		return nil, err
	}
	resp, _ := ioutil.ReadAll(cancel.Body)
	unmarshalError := json.Unmarshal(resp, &data)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	jsonStr, _ := json.Marshal(data[api.LastName(resource)])
	return jsonStr, err
}

func PatchId(sdkVersion string, host string, apiVersion string, language string, timeout int, user user.User, resource map[string]string, id string, payload interface{}) ([]byte, Error.StarkErrors) {
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
		nil,
	)
	if err.Errors != nil {
		return nil, err
	}
	resp, _ := ioutil.ReadAll(update.Body)
	unmarshalError := json.Unmarshal(resp, &data)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	jsonStr, _ := json.Marshal(data[api.LastName(resource)])
	return jsonStr, err
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