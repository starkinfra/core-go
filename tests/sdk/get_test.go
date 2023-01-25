package sdk

import (
	"fmt"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	Log "github.com/starkinfra/core-go/tests/utils/boleto/log"
	Event "github.com/starkinfra/core-go/tests/utils/event"
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestBoletoGet(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	boleto, err := Boleto.Get("4537841761648640")
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println(boleto.Id)
}

func TestBoletoQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 150

	boletos := Boleto.Query(params)
	for boleto := range boletos {
		fmt.Println("id:", boleto.Id)
	}
}

func TestBoletoPage(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 100

	boletos, cursor, err := Boleto.Page(params)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, boleto := range boletos {
		fmt.Println(boleto.Id)
	}
	fmt.Println(cursor)
}

func TestBoletoLogGet(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	boleto, err := Log.Get("4537841761648640")
	if err.Errors != nil {
		for _, err := range err.Errors {
			panic(err)
		}
	}
	fmt.Println(boleto)
}

func TestBoletoLogQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 300
	params["after"] = "2022-11-16"

	boletos := Log.Query(params)

	for boleto := range boletos {
		fmt.Println("i, boleto", boleto.Id)
	}
}

func TestBoletoLogPage(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 100
	params["after"] = "2022-11-16"

	boletos, cursor, err := Log.Page(params)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, boleto := range boletos {
		fmt.Println(boleto)
	}
	fmt.Println(cursor)
}

func TestPaymentGet(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	payment, err := Invoice.GetPayment("6543381610102784")
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println(payment)
}

func TestInvoiceGetPdf(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	invoice, err := Invoice.Pdf("6543381610102784")
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	filename := fmt.Sprintf("%v%v.pdf", api.Endpoint(Invoice.ResourceInvoice), "5767877416189952")
	createFileError := ioutil.WriteFile(filename, invoice, 0666)
	if createFileError != nil {
		fmt.Println(createFileError)
	}
}

func TestInvoiceGetQrcode(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	var paramsQrcode = map[string]interface{}{}
	paramsQrcode["size"] = 12

	invoice, err := Invoice.Qrcode("6543381610102784", paramsQrcode)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	filename := fmt.Sprintf("%v%v.png", api.Endpoint(Invoice.ResourceInvoice), "6543381610102784")
	createFileError := ioutil.WriteFile(filename, invoice, 0666)
	if createFileError != nil {
		fmt.Println(createFileError)
	}
}

func TestWorkspaceReplaceQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 1

	invoices := Invoice.Query(params, User.ExampleOrganization.Replace("4690697751887872"))
	for invoice := range invoices {
		fmt.Println("invoice's id: ", invoice.Id)
	}
}

func TestEventQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["isDelivered"] = true
	params["limit"] = 300

	events := Event.Query(params)

	for event := range events {
		fmt.Println(event.Id)
	}
}
