package sdk

import (
	"fmt"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
	"github.com/starkinfra/core-go/tests/utils/issuing/card"
	"github.com/starkinfra/core-go/tests/utils/issuing/holder"
	User "github.com/starkinfra/core-go/tests/utils/user"
	Webhook "github.com/starkinfra/core-go/tests/utils/webhook"
	"math/rand"
	"testing"
	"time"
)

func TestBoletoPostMulti(t *testing.T) {
	due := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	object := []Boleto.Boleto{
		{
			Amount:      rand.Intn(10000),
			Name:        "Core-Go-Test-multi-1",
			TaxId:       "38.446.231/0001-04",
			StreetLine1: "Kubasch Street, 900",
			StreetLine2: "OSLO",
			District:    "Ronny",
			City:        "Emmet City",
			StateCode:   "SP",
			ZipCode:     "01420-020",
			Due:         &due,
		},
	}

	boletos, err := Boleto.Create(object)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, boleto := range boletos {
		fmt.Println("boleto's id: ", boleto.Id)
		fmt.Println("boleto's amount: ", boleto.Amount)
	}
}

func TestInvoicePostMulti(t *testing.T) {
	invoiceToCreate := []Invoice.Invoice{
		{
			Amount: 111123,
			Name:   "Core-Go-Test-multi-1",
			TaxId:  "38.446.231/0001-04",
		}, {
			Amount: 222123,
			Name:   "Core-Go-Test-multi-1",
			TaxId:  "38.446.231/0001-04",
		},
	}

	invoices, err := Invoice.CreateWithUser(invoiceToCreate, User.ExampleOrganization.Replace("4690697751887872"))
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, invoice := range invoices {
		fmt.Println("invoice's id: ", invoice.Id)
		fmt.Println("invoice's amount: ", invoice.Amount)
	}
}

func TestWebhookPostSingle(t *testing.T) {
	object := Webhook.Webhook{
		Url:           "https://webhook.site/8d4903e7-cc3f-45f6-ac4c-6ce52abc4f0b",
		Subscriptions: []string{"boleto"},
	}

	webhook, err := Webhook.Create(object)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println("webhook's id: ", webhook.Id)
}

func TestIssuingHolderPost(t *testing.T) {

	holderExample := []holder.IssuingHolder{
		{
			Name:       "Irene Thompson",
			TaxId:      "52.792.530/0001-13",
			ExternalId: fmt.Sprintf("%v%v", time.Now().Format("20060102150405.999999999Z07001504"), "308764"),
		},
	}

	var expand = map[string]interface{}{}
	expand["expand"] = "rules"

	holders, err := holder.Create(holderExample, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, holder := range holders {
		fmt.Printf("%+v", holder)
	}
}

func TestIssuingCardPost(t *testing.T) {

	cardExample := []card.IssuingCard{
		{
			HolderName:       "Irene Thompson",
			HolderTaxId:      "52.792.530/0001-13",
			HolderExternalId: "308764",
		},
	}

	var expand = map[string]interface{}{}
	expand["expand"] = "securityCode, number, expiration"

	cards, err := card.Create(cardExample, expand)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, card := range cards {
		fmt.Printf("%+v\n", card)
	}
}
