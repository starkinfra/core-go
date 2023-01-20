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

	boletos, err := Boleto.Query(params)
	boleto, err := Boleto.Get(boletos[rand.Intn(params["limit"].(int))].Id)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println(boleto.Id)
}

func TestBoletoQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 3

	boletos, err := Boleto.Query(params)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, boleto := range boletos {
		fmt.Println(boleto)
	}
}

func TestBoletoPage(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 4

	boletos, cursor, err := Boleto.Page(params)
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

func TestBoletoLogGet(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	boletos, err := Log.Query(params)
	boleto, err := Log.Get(boletos[rand.Intn(params["limit"].(int))].Id)
	if err.Errors != nil {
		for _, err := range err.Errors {
			panic(err)
		}
	}
	fmt.Println(boleto)
}

func TestBoletoLogQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 108
	params["after"] = "2022-11-16"

	boletos, err := Log.Query(nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, boleto := range boletos {
		fmt.Println(boleto)
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

	invoices, err := Invoice.Query(params, User.ExampleProjectBank)
	payment, err := Invoice.GetPayment(invoices[rand.Intn(params["limit"].(int))].Id)
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

	invoices, err := Invoice.Query(params, User.ExampleProjectBank)
	invoice, err := Invoice.Pdf(invoices[rand.Intn(params["limit"].(int))].Id)
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

	invoices, errQuery := Invoice.Query(params, User.ExampleProjectBank)
	if errQuery.Errors != nil {
		for _, e := range errQuery.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	invoice, err := Invoice.Qrcode(invoices[rand.Intn(params["limit"].(int))].Id, paramsQrcode)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	filename := fmt.Sprintf("%v%v.png", api.Endpoint(Invoice.ResourceInvoice), invoices[rand.Intn(params["limit"].(int))].Id)
	createFileError := ioutil.WriteFile(filename, invoice, 0666)
	if createFileError != nil {
		fmt.Println(createFileError)
	}
}

func TestWorkspaceReplaceQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 1

	invoices, err := Invoice.Query(params, User.ExampleOrganization.Replace("4690697751887872"))
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, invoice := range invoices {
		fmt.Println("invoice's id: ", invoice.Id)
	}
}

func TestEventQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["isDelivered"] = true
	params["limit"] = 3

	events, err := Event.Query(params)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, event := range events {
		fmt.Println(event)
	}
}
