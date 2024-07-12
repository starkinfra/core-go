package sdk

import (
	"encoding/json"
	"fmt"
	"testing"
	"strconv"
	"time"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"github.com/stretchr/testify/assert"
)

func TestRawGet(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 2
	data := map[string]interface{}{}

	response, err := rest.GetRaw(
		"0.0.0",
		"bank",
		"v2",
		"pt-BR",
		15,
		"invoice",
		User.ExampleProjectBank,
		params,
		"Joker",
		false,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	fmt.Println(data)
	assert.NotNil(t, data)
}

func TestRawPost(t *testing.T) {
	body := map[string][]map[string]interface{}{
		"invoices": {
			{
				"amount": 996699999,
				"name":   "Tony Stark",
				"taxId":  "38.446.231/0001-04",
			},
		},
	}

	data := map[string]interface{}{}

	response, err := rest.PostRaw(
		"0.0.0",
		"bank",
		"v2",
		"pt-BR",
		15,
		"invoice",
		body,
		User.ExampleProjectBank,
		nil,
		"Joker",
		false,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	assert.NotNil(t, data)
}

func TestRawPatch(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 2
	data := map[string]interface{}{}

	response, err := rest.GetRaw(
		"0.0.0",
		"bank",
		"v2",
		"pt-BR",
		15,
		"invoice",
		User.ExampleProjectBank,
		params,
		"",
		false,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	invoicesData, ok1 := data["invoices"].([]interface{})
	if !ok1 {
		fmt.Println("Erro ao converter os tipos content")
		return
	}
	for _, invoice := range invoicesData {
		invoiceMap, _ := invoice.(map[string]interface{})
		id, _ := invoiceMap["id"].(string)
		path := "invoice/" + id
		for k := range data {
			delete(data, k)
		}
		body := map[string]interface{}{
			"amount": 0,
		}

		response, err := rest.PatchRaw(
			"0.0.0",
			"bank",
			"v2",
			"pt-BR",
			15,
			path,
			body,
			User.ExampleProjectBank,
			nil,
			"Joker",
			false,
		)

		if err.Errors != nil {
			for _, e := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
			}
		}
		unmarshalError := json.Unmarshal(response.Content, &data)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		invoiceData, _ := data["invoice"].(map[string]interface{})
		amount, _ := invoiceData["amount"].(int)
		assert.Equal(t, 0, amount)
	}
}

func TestRawPut(t *testing.T) {
	data := map[string]interface{}{}
	body := map[string][]map[string]interface{}{
		"profiles": {
			{
				"interval": "day",
				"delay":    0,
			},
		},
	}
	path := "split-profile/"

	response, err := rest.PutRaw(
		"0.0.0",
		"bank",
		"v2",
		"pt-BR",
		15,
		path,
		body,
		User.ExampleProjectBank,
		nil,
		"Joker",
		false,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	assert.NotNil(t, data)
}

func TestRequestDelete(t *testing.T) {
	data := map[string]interface{}{}

	now := time.Now()
	futureDate := now.AddDate(0, 0, 10).Format("2006-01-02")
	milliseconds := now.UnixNano() / int64(time.Millisecond)
	timestamp := strconv.FormatInt(milliseconds, 10)

	body := map[string][]map[string]interface{}{
		"transfers": {
			{
				"amount":        10000,
				"name":          "Steve Rogers",
				"taxId":         "330.731.970-10",
				"bankCode":      "001",
				"branchCode":    "1234",
				"accountNumber": "123456-0",
				"accountType":   "checking",
				"scheduled":     futureDate,
				"externalId":    timestamp,
			},
		},
	}

	path := "transfer/"
	response, err := rest.PostRaw(
		"0.0.0",
		"bank",
		"v2",
		"pt-BR",
		15,
		path,
		body,
		User.ExampleProjectBank,
		nil,
		"Joker",
		false,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	transfersData, ok1 := data["transfers"].([]interface{})
	if !ok1 {
		fmt.Println("Erro ao converter os tipos content")
		return
	}
	for _, transfer := range transfersData {
		transferMap, ok2 := transfer.(map[string]interface{})
		if !ok2 {
			fmt.Println("Erro ao converter item de list 'invoices' para map[string]interface{}")
			continue
		}
		id, ok3 := transferMap["id"].(string)
		if !ok3 {
			fmt.Println("Erro ao converter list 'id' para string")
			continue
		}
		path = "transfer/" + id
		for k := range data {
			delete(data, k)
		}
		response, err := rest.DeleteRaw(
			"0.0.0",
			"bank",
			"v2",
			"pt-BR",
			15,
			path,
			User.ExampleProjectBank,
			"",
			false,
		)

		if err.Errors != nil {
			for _, e := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
			}
		}
		unmarshalError := json.Unmarshal(response.Content, &data)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		assert.NotNil(t, data)
	}
}
