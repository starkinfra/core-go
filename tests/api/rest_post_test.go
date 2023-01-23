package api

import (
	"encoding/json"
	"fmt"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
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
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceBoleto,
		object,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(create, &boletos)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}

	for _, boleto := range boletos {
		fmt.Println(boleto)
	}
}
