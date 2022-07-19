package test

import (
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/hosts"
	"core-go/starkcore/utils/rest"
	"core-go/tests/utils/user"
	"testing"
)

type boletos struct {
	Boleto []boleto
}

type boleto struct {
	Amount        *int
	Name          *string
	TaxId         *string
	StreetLine1   *string
	StreetLine2   *string
	District      *string
	City          *string
	StateCode     *string
	ZipCode       *string
	Due           *string
	Fine          *float32
	Interest      *float32
	OverdueLimit  *int
	Descriptions  *map[string]string
	Discounts     *map[string]string
	Tags          *[]string
	ReceiverName  *string
	ReceiverTaxId *string
	Id            *string
	Fee           *int
	Line          *string
	BarCode       *string
	Transactions  *[]string
	Created       *string
	OurNumber     *string
}

var resource_s = map[string]string{"name": "boleto"}

func TestSuccessSingle(t *testing.T) {

	var amount = 1111111111
	var name = "PRIMEIRO TESTE DE POST SINGLE NA TERÇA FEIRA"
	var taxId = "38.446.231/0001-04"
	var streetLine1 = "Kubasch Street, 900"
	var streetLine2 = ""
	var district = "Ronny"
	var city = "Emmet City"
	var stateCode = "SP"
	var zipCode = "01420-020"

	var boletoSingle = []boleto{
		{
			Amount:      &amount,
			Name:        &name,
			TaxId:       &taxId,
			StreetLine1: &streetLine1,
			StreetLine2: &streetLine2,
			District:    &district,
			City:        &city,
			StateCode:   &stateCode,
			ZipCode:     &zipCode,
		},
	}

	var boletosingle = boletos{Boleto: boletoSingle}

	rest.PostSingle(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_g,
		api.ApiJson(boletosingle, resource_s),
		"pt-BR",
	)
}

func TestSuccessPostMulti(t *testing.T) {

	var amount2 = 22222222222
	var name2 = "PRIMEIRO TESTE DE POSTMULTI  NA TERÇA FEIRA"
	var taxId2 = "38.446.231/0001-04"
	var streetLine12 = "Kubasch Street, 900"
	var streetLine22 = ""
	var district2 = "Ronny"
	var city2 = "Emmet City"
	var stateCode2 = "SP"
	var zipCode2 = "01420-020"

	var amount3 = 33333333333
	var name3 = "SEGUNDO TESTE DE POSTMULTI  NA TERÇA FEIRA"
	var taxId3 = "38.446.231/0001-04"
	var streetLine13 = "Kubasch Street, 900"
	var streetLine23 = ""
	var district3 = "Ronny"
	var city3 = "Emmet City"
	var stateCode3 = "SP"
	var zipCode3 = "01420-020"

	var boletoMulti = []boleto{
		{
			Amount:      &amount2,
			Name:        &name2,
			TaxId:       &taxId2,
			StreetLine1: &streetLine12,
			StreetLine2: &streetLine22,
			District:    &district2,
			City:        &city2,
			StateCode:   &stateCode2,
			ZipCode:     &zipCode2,
		},
		{
			Amount:      &amount3,
			Name:        &name3,
			TaxId:       &taxId3,
			StreetLine1: &streetLine13,
			StreetLine2: &streetLine23,
			District:    &district3,
			City:        &city3,
			StateCode:   &stateCode3,
			ZipCode:     &zipCode3,
		},
	}

	var boletomulti = boletos{Boleto: boletoMulti}

	rest.PostMulti(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_g,
		api.ApiJson(boletomulti, resource_g),
		"pt-BR",
	)
}
