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
	Amount        int
	Name          string
	TaxId         string
	StreetLine1   string
	District      string
	StreetLine2   string
	City          string
	StateCode     string
	ZipCode       string
	Due           string
	Fine          float32
	Interest      float32
	OverdueLimit  int
	Descriptions  map[string]string
	Discounts     map[string]string
	Tags          []string
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

var resource_p = map[string]string{"name": "boleto"}

func TestSuccessSingle(t *testing.T) {

	receiverName := "Dotts"
	var boletop = []boleto{
		{
			Amount:       123456789,
			Name:         "PRIMEIRO TESTE NOVO NA SEGUNDA-FEIRA",
			TaxId:        "38.446.231/0001-04",
			StreetLine1:  "Kubasch Street, 900",
			StreetLine2:  "",
			District:     "Ronny",
			City:         "Emmet City",
			StateCode:    "SP",
			ZipCode:      "01420-020",
			ReceiverName: &receiverName,
		},
		{
			Amount:      123456789,
			Name:        "SEGUNDO TESTE NOVO NA SEGUNDA-FEIRA",
			TaxId:       "38.446.231/0001-04",
			StreetLine1: "Kubasch Street, 900",
			StreetLine2: "",
			District:    "Ronny",
			City:        "Emmet City",
			StateCode:   "SP",
			ZipCode:     "01420-020",
		},
	}

	var boletos_example = boletos{Boleto: boletop}

	rest.PostSingle(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_g,
		api.ApiJson(boletos_example, resource_g),
		"pt-BR",
	)
}
