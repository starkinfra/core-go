package sdk

import (
	"fmt"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
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

//func TestInvoicePostMulti(t *testing.T) {
//	invoiceToCreate := []Invoice.Invoice{
//		{
//			Amount: 111123,
//			Name:   "Core-Go-Test-multi-1",
//			TaxId:  "38.446.231/0001-04",
//		}, {
//			Amount: 222123,
//			Name:   "Core-Go-Test-multi-1",
//			TaxId:  "38.446.231/0001-04",
//		},
//	}
//
//	invoices, err := Invoice.CreateWithUser(invoiceToCreate, User.ExampleOrganization.Replace("4690697751887872"))
//	if err.Errors != nil {
//		for _, e := range err.Errors {
//			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
//		}
//	}
//	for _, invoice := range invoices {
//		fmt.Println("invoice's id: ", invoice.Id)
//		fmt.Println("invoice's amount: ", invoice.Amount)
//	}
//}

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
