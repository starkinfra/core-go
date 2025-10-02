package examples

import (
	"math/rand"
	"strconv"
	"time"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
)

func ExampleBoleto() []Boleto.Boleto {
	due := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	boleto := []Boleto.Boleto{
		{
			Amount:      rand.Intn(100000),
			Name:        "Core-Go-Test-Boleto",
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
	return boleto
}

func ExampleInvoice() []Invoice.Invoice {
	invoice := []Invoice.Invoice{
		{
			Name: "Arya Stark",
			TaxId: "38.446.231/0001-04",
			Amount: 10000,
		},
	}
	return invoice
}

func ExampleTransferBody() map[string][]map[string]interface{} {
	now := time.Now()
	futureDate := now.AddDate(0, 0, 10).Format("2006-01-02")
	milliseconds := now.UnixNano() / int64(time.Millisecond)
	timestamp := strconv.FormatInt(milliseconds, 10)

	transferBody := map[string][]map[string]interface{}{
		"transfers": {
			{
				"amount":        10000,
				"name":          "Steve Rogers",
				"taxId":         "073.310.980-20",
				"bankCode":      "001",
				"branchCode":    "1234",
				"accountNumber": "123456-0",
				"accountType":   "checking",
				"scheduled":     futureDate,
				"externalId":    timestamp,
			},
		},
	}
	return transferBody
}