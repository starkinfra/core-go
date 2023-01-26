package card

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"time"
)

type IssuingCard struct {
	HolderName       string        `json:",omitempty"`
	HolderTaxId      string        `json:",omitempty"`
	HolderExternalId string        `json:",omitempty"`
	DisplayName      string        `json:",omitempty"`
	Rules            []IssuingRule `json:",omitempty"`
	ProductId        string        `json:",omitempty"`
	Tags             []string      `json:",omitempty"`
	StreetLine1      string        `json:",omitempty"`
	StreetLine2      string        `json:",omitempty"`
	District         string        `json:",omitempty"`
	City             string        `json:",omitempty"`
	StateCode        string        `json:",omitempty"`
	ZipCode          string        `json:",omitempty"`
	Id               string        `json:",omitempty"`
	HolderId         string        `json:",omitempty"`
	Type             string        `json:",omitempty"`
	Status           string        `json:",omitempty"`
	Number           string        `json:",omitempty"`
	SecurityCode     string        `json:",omitempty"`
	Expiration       string        `json:",omitempty"`
	Updated          *time.Time    `json:",omitempty"`
	Created          *time.Time    `json:",omitempty"`
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

var objects []IssuingCard
var resourceIssuingCard = map[string]string{"name": "IssuingCard"}

func Create(cards []IssuingCard, expand map[string]interface{}) ([]IssuingCard, Error.StarkErrors) {
	create, err := rest.PostMulti(
		utils.SdkVersion,
		hosts.Infra,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectInfra,
		resourceIssuingCard,
		cards,
		expand,
	)
	unmarshalError := json.Unmarshal(create, &objects)
	if unmarshalError != nil {
		return objects, err
	}
	return objects, err
}
