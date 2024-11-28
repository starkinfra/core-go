package holder

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"time"
)

type IssuingHolder struct {
	Name       string        `json:",omitempty"`
	TaxId      string        `json:",omitempty"`
	ExternalId string        `json:",omitempty"`
	Rules      []IssuingRule `json:",omitempty"`
	Tags       []string      `json:",omitempty"`
	Id         string        `json:",omitempty"`
	Status     string        `json:",omitempty"`
	Updated    *time.Time    `json:",omitempty"`
	Created    *time.Time    `json:",omitempty"`
}

type IssuingRule struct {
	Name           string             `json:",omitempty"`
	Amount         int                `json:",omitempty"`
	Id             string             `json:",omitempty"`
	Interval       string             `json:",omitempty"`
	CurrencyCode   string             `json:",omitempty"`
	Categories     []MerchantCategory `json:",omitempty"`
	Countries      []MerchantCountry  `json:",omitempty"`
	Methods        []CardMethod       `json:",omitempty"`
	CounterAmount  int                `json:",omitempty"`
	CurrencySymbol string             `json:",omitempty"`
	CurrencyName   string             `json:",omitempty"`
}

type MerchantCategory struct {
	Code   string `json:",omitempty"`
	Type   string `json:",omitempty"`
	Name   string `json:",omitempty"`
	Number string `json:",omitempty"`
}

type MerchantCountry struct {
	Code      string `json:",omitempty"`
	Name      string `json:",omitempty"`
	Number    string `json:",omitempty"`
	ShortCode string `json:",omitempty"`
}

type CardMethod struct {
	Code   string `json:",omitempty"`
	Name   string `json:",omitempty"`
	Number string `json:",omitempty"`
}

var resourceIssuingHolder = map[string]string{"name": "IssuingHolder"}

func Create(holders []IssuingHolder, expand map[string]interface{}) ([]IssuingHolder, Error.StarkErrors) {
	var objects []IssuingHolder
	create, err := rest.PostMulti(
		utils.SdkVersion,
		hosts.Infra,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectInfra,
		resourceIssuingHolder,
		holders,
		expand,
	)
	unmarshalError := json.Unmarshal(create, &objects)
	if unmarshalError != nil {
		return objects, err
	}
	return objects, err
}
