package test

import (
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/hosts"
	"core-go/starkcore/utils/rest"
	"core-go/tests/utils/user"
	"testing"
)

type invoice struct {
	Amount         *int
	Name           *string
	TaxId          *string
	Due            *string
	Fine           *float32
	Interest       *float32
	Expiration     *int
	Descriptions   *map[string]string
	Discounts      *map[string]string
	Tags           *[]string
	Pdf            *string
	Link           *string
	NominalAmount  *int
	FineAmount     *int
	InterestAmount *int
	DiscountAmount *int
	Id             *string
	Brcode         *string
	Status         *string
	Fee            *int
	TransactionIds *[]string
	Created        *string
	Updated        *string
}

var resource_v = map[string]string{"name": "invoice"}

func TestSuccessPatch(t *testing.T) {

	status := "canceled"

	invoicePatch := invoice{
		Status: &status,
	}

	rest.PatchId(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_v,
		"6459912880128000",
		api.ApiJson(invoicePatch),
		"pt-BR",
		15,
		nil,
	)
}
