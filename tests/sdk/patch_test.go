package sdk

import (
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
	"github.com/starkinfra/core-go/tests/utils/examples"
	"testing"
)

func TestInvoicePostAndPatch(t *testing.T) {

	object := examples.ExampleInvoice()

	invoicesCreated, errCreate := Invoice.Create(object)
	if errCreate.Errors != nil {
		for _, e := range errCreate.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	invoiceUpdated, err := Invoice.Update(invoicesCreated[0].Id)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	if invoiceUpdated.Id == "" {
		t.Errorf("invoiceUpdated.Id is empty")
	}
}
