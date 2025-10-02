package sdk

import (
	"fmt"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	Log "github.com/starkinfra/core-go/tests/utils/boleto/log"
	Event "github.com/starkinfra/core-go/tests/utils/event"
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
	"github.com/starkinfra/core-go/tests/utils/issuing/product"
	"github.com/starkinfra/core-go/tests/utils/sign"
	"github.com/starkinfra/core-go/tests/utils/examples"
	"os"
	"testing"
)

func TestBoletoCreateAndGet(t *testing.T) {
	createdBoleto, err := Boleto.Create(examples.ExampleBoleto())
	if err.Errors != nil {
		t.Errorf("err: %s", err.Errors)
	}

	_, err = Boleto.Get(createdBoleto[0].Id)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
}

func TestBoletoQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 200

	boletos, err := Boleto.Query(params)

	go func() {
		for err := range err {
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		}
	}()

	for boleto := range boletos {
		if boleto.Id == "" {
			t.Errorf("boleto.Id is empty")
		}
	}
}

func TestBoletoPage(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 100

	boletos, _, err := Boleto.Page(params)
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

func TestCreateBoletoAndLogGet(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 1

	_, err := Boleto.Create(examples.ExampleBoleto())

	if err.Errors != nil {
		t.Errorf("err: %s", err.Errors)
	}

	boletoLogs, _, err := Log.Page(params)

	boletoLogId := boletoLogs[0].Id

	_, err = Log.Get(boletoLogId)

	if err.Errors != nil {
		for _, err := range err.Errors {
			t.Errorf("code: %s, message: %s", err.Code, err.Message)
		}
	}
}

func TestDocumentGet(t *testing.T) {
	_, err := sign.Get("52e2ab8389dd4fa5856b095ce6a9b125")
	if err.Errors != nil {
		for _, err := range err.Errors {
			panic(err)
		}
	}
}

func TestBoletoLogQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 300
	params["after"] = "2022-11-16"

	boletos, err := Log.Query(params)

	go func() {
		for err := range err {
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		}
	}()

	for boleto := range boletos {
		if boleto.Id == "" {
			t.Errorf("boleto.Id is empty")
		}
	}
}

func TestProductLogQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 100

	products, err := product.Query(params)

	go func() {
		for err := range err {
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		}
	}()

	for product := range products {
		if product.Id == "" {
			t.Errorf("product.Id is empty")
		}
	}
}

func TestBoletoLogPage(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 100
	params["after"] = "2022-11-16"

	boletos, _, err := Log.Page(params)
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

func TestInvoicePageAndGetPayment(t *testing.T) {
	var params = map[string]interface{}{}
	params["status"] = "paid"
	params["limit"] = 1

	payments, _, err := Invoice.Page(params)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	payment, err := Invoice.GetPayment(payments[0].Id)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	if payment.EndToEndId == "" {
		t.Errorf("payment.Id is empty")
	}
}

func TestInvoiceGetPdf(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 1

	invoices, _, err := Invoice.Page(params)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	invoice, err := Invoice.Pdf(invoices[0].Id)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	filename := fmt.Sprintf("%v%v.pdf", api.Endpoint(Invoice.ResourceInvoice), invoices[0].Id)
	createFileError := os.WriteFile(filename, invoice, 0666)
	if createFileError != nil {
		t.Errorf("createFileError: %s", createFileError)
	}
}

func TestInvoiceGetQrcode(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 1

	var paramsQrcode = map[string]interface{}{}
	paramsQrcode["size"] = 12

	invoices, _, err := Invoice.Page(params)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	invoice, err := Invoice.Qrcode(invoices[0].Id, paramsQrcode)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	filename := fmt.Sprintf("%v%v.png", api.Endpoint(Invoice.ResourceInvoice), invoices[0].Id)
	createFileError := os.WriteFile(filename, invoice, 0666)
	if createFileError != nil {
		t.Errorf("createFileError: %s", createFileError)
	}
}

func TestEventQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["isDelivered"] = true
	params["limit"] = 300

	events, err := Event.Query(params)

	go func() {
		for err := range err {
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		}
	}()

	for event := range events {
		if event.Id == "" {
			t.Errorf("event.Id is empty")
		}
	}
}
