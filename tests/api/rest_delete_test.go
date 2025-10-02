package api

import (
	"encoding/json"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"github.com/starkinfra/core-go/tests/utils/examples"
	"testing"
)

func TestSuccessDelete(t *testing.T) {
	var boleto Boleto.Boleto
	var newBoleto []Boleto.Boleto

	object := examples.ExampleBoleto()
	create, err := rest.PostMulti(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoleto,
		object,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(create, &newBoleto)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
	}

	var id string
	for _, boleto := range newBoleto {
		id = boleto.Id
	}

	deleted, err := rest.DeleteId(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoleto,
		id,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError = json.Unmarshal(deleted, &boleto)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
	}
}

func TestFailDelete(t *testing.T) {
	_, err := rest.DeleteId(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoleto,
		"6678130383126528",
		nil,
	)
	if err.Errors != nil {
		if err.Errors[0].Code != "invalidBoleto" {
			t.Errorf("err.Errors: %+v\n", err.Errors)
		}
		return
	}
	t.Errorf("Test failed")
}