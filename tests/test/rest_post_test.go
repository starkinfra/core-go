package test

import (
	"core-go/starkcore/utils/api"
	"core-go/starkcore/utils/hosts"
	"core-go/starkcore/utils/rest"
	"core-go/tests/utils/user"
	"fmt"
	"testing"
)

type boletos struct {
	Boleto []boleto `json:"boletos"`
}

type boleto struct {
	Amount        int               `json:"amount"`
	Name          string            `json:"name"`
	TaxId         string            `json:"taxId"`
	StreetLine1   string            `json:"streetLine1"`
	StreetLine2   string            `json:"streetLine2"`
	District      string            `json:"district"`
	City          string            `json:"city"`
	StateCode     string            `json:"stateCode"`
	ZipCode       string            `json:"zipCode"`
	Due           string            `json:"due"`
	Fine          float32           `json:"fine"`
	Interest      float32           `json:"interest"`
	OverdueLimit  int               `json:"overdueLimit"`
	Descriptions  map[string]string `json:"descriptions"`
	Discounts     map[string]string `json:"discounts"`
	Tags          []string          `json:"tags"`
	ReceiverName  string            `json:"receiverName"`
	ReceiverTaxId string            `json:"receiverTaxId"`
	Id            string            `json:"id"`
	Fee           int               `json:"fee"`
	Line          string            `json:"line"`
	BarCode       string            `json:"barCode"`
	Transactions  []string          `json:"transactions"`
	Created       string            `json:"created"`
	OurNumber     string            `json:"ourNumber"`
}

var resource_p = map[string]string{"name": "boleto"}

func TestSuccessSingle(t *testing.T) {

	var boletop = []boleto{
		{
			Amount:        61807,
			Name:          "Alec Feuerborn",
			TaxId:         "38.446.231/0001-04",
			StreetLine1:   "Kubasch Street, 900",
			StreetLine2:   "",
			District:      "Ronny",
			City:          "Emmet City",
			StateCode:     "SP",
			ZipCode:       "01420-020",
			Due:           "2022-07-13",
			Fine:          2,
			Interest:      1,
			OverdueLimit:  59,
			Descriptions:  nil,
			Discounts:     nil,
			Tags:          nil,
			ReceiverName:  "Stark Bank S.A. - Instituicao de Pagamento",
			ReceiverTaxId: "20.018.183/0001-80",
		},
	}

	var boletos_example = boletos{Boleto: boletop}

	var transactions, _ = rest.PostSingle(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_g,
		api.ApiJson(boletos_example),
		"pt-BR",
	)

	fmt.Printf("\nTRANSACTIONS:::%v\n", transactions)
}

//
//func TestSuccessMulti(t *testing.T) {
//	transactions, _ := rest.PostMulti(
//		"0.0.0",
//		hosts.Service.Bank,
//		"v2",
//		user.ExampleProject,
//		resource_g,
//		&exampleBoleto,
//		"pt-BR",
//	)
//
//	fmt.Sprintf("%v", transactions)
//}

