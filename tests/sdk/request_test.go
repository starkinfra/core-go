package sdk

import (
	"encoding/json"
	"testing"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils/examples"
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
		true,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
	}
	assert.NotNil(t, data)
}

func TestRawPost(t *testing.T) {
	body := examples.ExampleInvoice()

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
		true,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
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
		true,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
	}
	invoicesData, conversionError := data["invoices"].([]interface{})
	if !conversionError {
		t.Errorf("Erro ao converter os tipos content")
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
			true,
		)

		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		unmarshalError := json.Unmarshal(response.Content, &data)
		if unmarshalError != nil {
			t.Errorf("unmarshalError: %s", unmarshalError)
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
		true,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
	}
	assert.NotNil(t, data)
}

func TestRequestDelete(t *testing.T) {
	data := map[string]interface{}{}

	body := examples.ExampleTransferBody()

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
		true,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
	}

	transfersData, conversionError := data["transfers"].([]interface{})
	if !conversionError {
		t.Errorf("Erro ao converter os tipos content")
		return
	}
	for _, transfer := range transfersData {
		transferMap, conversionError := transfer.(map[string]interface{})
		if !conversionError {
			t.Errorf("Erro ao converter item de list 'invoices' para map[string]interface{}")
			continue
		}
		id, conversionError := transferMap["id"].(string)
		if !conversionError {
			t.Errorf("Erro ao converter list 'id' para string")
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
			true,
		)

		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		unmarshalError := json.Unmarshal(response.Content, &data)
		if unmarshalError != nil {
			t.Errorf("unmarshalError: %s", unmarshalError)
		}
		assert.NotNil(t, data)
	}
}
