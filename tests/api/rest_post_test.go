package api

import (
	"encoding/json"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	"github.com/starkinfra/core-go/tests/utils/examples"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"math/rand"
	"testing"
	"time"
)

func TestSuccessPostMulti(t *testing.T) {
	var boletos []Boleto.Boleto
	due := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	object := []Boleto.Boleto{
		{
			Amount:      rand.Intn(10000),
			Name:        "Core-Go-Test-multi-1",
			TaxId:       "38.446.231/0001-04",
			StreetLine1: "Kubasch Street, 900",
			StreetLine2: "jhvh",
			District:    "Ronny",
			City:        "Emmet City",
			StateCode:   "SP",
			ZipCode:     "01420-020",
			Due:         &due,
		},
		{
			Amount:      rand.Intn(10000),
			Name:        "Core-Go-Test-multi-1",
			TaxId:       "38.446.231/0001-04",
			StreetLine1: "Kubasch Street, 900",
			StreetLine2: "jhvh",
			District:    "Ronny",
			City:        "Emmet City",
			StateCode:   "SP",
			ZipCode:     "01420-020",
			Due:         &due,
		},
	}
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
	unmarshalError := json.Unmarshal(create, &boletos)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
	}

	for _, boleto := range boletos {
		if boleto.Id == "" {
			t.Errorf("boleto.Id is empty")
		}
	}
}

func TestFailPostMulti(t *testing.T) {
	object := examples.ExampleBoleto()
	object[0].Amount = 0

	_, err := rest.PostMulti(
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
		return
	}
	t.Errorf("Test failed")
}