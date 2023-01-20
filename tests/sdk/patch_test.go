package sdk

import (
	"fmt"
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
	"math/rand"
	"testing"
)

func TestInvoicePostAndPatch(t *testing.T) {

	object := []Invoice.Invoice{
		{
			Amount: rand.Intn(10000),
			Name:   "Core-Go-Test-multi-1",
			TaxId:  "38.446.231/0001-04",
		},
	}

	invoicesCreated, errCreate := Invoice.Create(object)
	if errCreate.Errors != nil {
		for _, e := range errCreate.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println("invoice's id: ", invoicesCreated[0].Id)
	fmt.Println("invoice's amount: ", invoicesCreated[0].Amount)

	invoiceUpdated, err := Invoice.Update(invoicesCreated[0].Id)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println("invoice's id", invoiceUpdated.Id)
	fmt.Println("invoice's amount: ", invoiceUpdated.Amount)
}
