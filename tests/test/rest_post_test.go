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
	Amount        int
	Name          string
	TaxId         string
	StreetLine1   string
	StreetLine2   string
	District      string
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
	ReceiverName  string
	ReceiverTaxId string
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

	var boletop = []boleto{
		{
			Amount:      123456789,
			Name:        "TESTE NOVO SEXTA-FEIRA",
			TaxId:       "38.446.231/0001-04",
			StreetLine1: "Kubasch Street, 900",
			StreetLine2: "",
			District:    "Ronny",
			City:        "Emmet City",
			StateCode:   "SP",
			ZipCode:     "01420-020",
		},
		{
			Amount:      987654321,
			Name:        "SEGUNDO TESTE DE SEXTA-FEIRA",
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
	//fmt.Println(boletos_example)
	//
	//var m = make(map[string]interface{})
	//
	//out, _ := json.Marshal(boletos_example)
	//json.Unmarshal(out, &m)
	//fmt.Println(string(out))
	//fmt.Println(m)

	//for key, value := range m {
	//	if value != "" {
	//		fmt.Printf("%v", value)
	//		delete(m, key)
	//	}
	//	fmt.Printf("%v:%v", key, value)

	//if err != nil {
	//	return
	//}

	//jsons, err := json.Marshal(CastJsonToApiFormat(m))
	//if err != nil {
	//	panic(err)
	//}
	//
	//return string(jsons)

	var transactions, _ = rest.PostSingle(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_g,
		api.ApiJson(boletos_example),
		"pt-BR",
	)

	fmt.Println(transactions)
}
