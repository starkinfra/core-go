package sdk

import (
	"fmt"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
	"github.com/starkinfra/core-go/tests/utils/issuing/card"
	"github.com/starkinfra/core-go/tests/utils/issuing/holder"
	User "github.com/starkinfra/core-go/tests/utils/user"
	Webhook "github.com/starkinfra/core-go/tests/utils/webhook"
	"github.com/starkinfra/core-go/tests/utils/examples"
	"strconv"
	"testing"
	"time"
)

func TestBoletoPostMulti(t *testing.T) {
	object := examples.ExampleBoleto()

	boletos, err := Boleto.Create(object)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, boleto := range boletos {
		if boleto.Id == "" {
			t.Errorf("boleto.Id is empty")
		}
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

	invoices, err := Invoice.CreateWithUser(invoiceToCreate, User.ExampleProjectBank)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, invoice := range invoices {
		if invoice.Id == "" {
			t.Errorf("invoice.Id is empty")
		}
	}
}

func TestWebhookPostSingle(t *testing.T) {
	milliseconds := time.Now().UnixNano() / int64(time.Millisecond)
	timestamp := strconv.FormatInt(milliseconds, 10)
	object := Webhook.Webhook{
		Url:           "https://webhook.site/" + timestamp,
		Subscriptions: []string{"boleto"},
	}

	webhook, err := Webhook.Create(object)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	if webhook.Id == "" {
		t.Errorf("webhook.Id is empty")
	}
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
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, holder := range holders {
		if holder.Id == "" {
			t.Errorf("holder.Id is empty")
		}
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
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, card := range cards {
		if card.Id == "" {
			t.Errorf("card.Id is empty")
		}
	}
}
