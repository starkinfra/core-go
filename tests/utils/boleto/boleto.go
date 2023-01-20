package boleto

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"time"
)

type Boleto struct {
	Amount        int                      `json:",omitempty"`
	Name          string                   `json:",omitempty"`
	TaxId         string                   `json:",omitempty"`
	StreetLine1   string                   `json:",omitempty"`
	StreetLine2   string                   `json:",omitempty"`
	District      string                   `json:",omitempty"`
	City          string                   `json:",omitempty"`
	StateCode     string                   `json:",omitempty"`
	ZipCode       string                   `json:",omitempty"`
	Due           *time.Time               `json:",omitempty"`
	Fine          float64                  `json:",omitempty"`
	Interest      float64                  `json:",omitempty"`
	OverdueLimit  int                      `json:",omitempty"`
	Descriptions  []map[string]interface{} `json:",omitempty"`
	Discounts     []map[string]interface{} `json:",omitempty"`
	Tags          []string                 `json:",omitempty"`
	ReceiverName  string                   `json:",omitempty"`
	ReceiverTaxId string                   `json:",omitempty"`
	Id            string                   `json:",omitempty"`
	Fee           int                      `json:",omitempty"`
	Line          string                   `json:",omitempty"`
	BarCode       string                   `json:",omitempty"`
	Transactions  []string                 `json:",omitempty"`
	Created       *time.Time               `json:",omitempty"`
	OurNumber     string                   `json:",omitempty"`
}

var resourceBoleto = map[string]string{"name": "Boleto"}
var boletos []Boleto
var boleto Boleto

func Create(boletos []Boleto) ([]Boleto, Error.StarkErrors) {
	create, err := rest.PostMulti(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		resourceBoleto,
		boletos,
	)
	if err.Errors != nil {
		return []Boleto{}, err
	}
	unmarshalError := json.Unmarshal(create, &boletos)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return boletos, err
}

func Get(id string) (Boleto, Error.StarkErrors) {
	get, err := rest.GetId(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		resourceBoleto,
		id,
	)
	if err.Errors != nil {
		return Boleto{}, err
	}
	unmarshalError := json.Unmarshal(get, &boleto)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return boleto, err
}

func Query(params map[string]interface{}) ([]Boleto, Error.StarkErrors) {
	query, err := rest.GetStream(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		resourceBoleto,
		params,
	)
	if err.Errors != nil {
		return []Boleto{}, err
	}
	unmarshalError := json.Unmarshal(query, &boletos)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return boletos, err
}

func Page(params map[string]interface{}) ([]Boleto, string, Error.StarkErrors) {
	page, cursor, err := rest.GetPage(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		resourceBoleto,
		params,
	)
	if err.Errors != nil {
		return []Boleto{}, "", err
	}
	unmarshalError := json.Unmarshal(page, &boletos)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return boletos, cursor, err
}

func Cancel(id string) (Boleto, Error.StarkErrors) {
	cancel, err := rest.DeleteId(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		resourceBoleto,
		id,
	)
	if err.Errors != nil {
		return Boleto{}, err
	}
	unmarshalError := json.Unmarshal(cancel, &boleto)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return boleto, err
}
