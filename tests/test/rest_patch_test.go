package test

import (
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/rest"
	"core-go/tests/utils/user"
	"testing"
)

type invoices struct {
	Invoices []invoice
}

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

	var status = "canceled"

	var invoicePatch = invoice{
		Status: &status,
	}

	rest.PatchId(
		"0.0.0",
		"",
		"v2",
		user.ExampleProject,
		resource_v,
		"",
		api.ApiJson(invoicePatch),
		"pt-BR",
		15,
	)
}
