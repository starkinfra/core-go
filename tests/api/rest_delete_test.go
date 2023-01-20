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
)

func TestSuccessDel(t *testing.T) {
	var boleto Boleto.Boleto
	var boletos []Boleto.Boleto

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	query, err := rest.GetStream(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceBoleto,
		params,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	queryError := json.Unmarshal(query, &boletos)
	if queryError != nil {
		fmt.Println(queryError)
	}

	deleted, err := rest.DeleteId(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceBoleto,
		boletos[rand.Intn(params["limit"].(int))].Id,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(deleted, &boleto)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	fmt.Println(boleto.Id)
}
