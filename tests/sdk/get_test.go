package sdk

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/starkinfra/core-go/starkcore/utils/api"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	Log "github.com/starkinfra/core-go/tests/utils/boleto/log"
	Event "github.com/starkinfra/core-go/tests/utils/event"
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
	"github.com/starkinfra/core-go/tests/utils/issuing/product"
	"github.com/starkinfra/core-go/tests/utils/sign"
	User "github.com/starkinfra/core-go/tests/utils/user"
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
	params["limit"] = 200

	boletos, err := Boleto.Query(params)

	for {
		select {
		case errors := <-err:
			if errors != nil {
				panic("your custom panic here")
			}
			return
		case result := <-boletos:
			fmt.Println(result.Id)
		}
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

func TestDocumentGet(t *testing.T) {
	document, err := sign.Get("52e2ab8389dd4fa5856b095ce6a9b125")
	if err.Errors != nil {
		for _, err := range err.Errors {
			panic(err)
		}
	}
	fmt.Println(document)
}

func TestBoletoLogQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 300
	params["after"] = "2022-11-16"

	boletos, err := Log.Query(params)

	for {
		select {
		case errors := <-err:
			if errors != nil {
				panic("your custom panic here")
			}
			return
		case result := <-boletos:
			fmt.Println(result.Id)
		}
	}
}

func TestProductLogQuery(t *testing.T) {

	products, err := product.Query(nil)

	for {
		select {
		case errors := <-err:
			if errors != nil {
				panic("your custom panic here")
			}
			return
		case result := <-products:
			fmt.Println(result.Id)
		}
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

	invoices, err := Invoice.Query(params, User.ExampleOrganization.Replace("4690697751887872"))
	for {
		select {
		case errors := <-err:
			if errors != nil {
				panic("your custom panic here")
			}
			return
		case result := <-invoices:
			fmt.Println(result.Id)
		}
	}
}

func TestEventQuery(t *testing.T) {
	var params = map[string]interface{}{}
	params["isDelivered"] = true
	params["limit"] = 300

	events, err := Event.Query(params)
	for {
		select {
		case errors := <-err:
			if errors != nil {
				panic("your custom panic here")
			}
			return
		case result := <-events:
			fmt.Println(result.Id)
		}
	}
}
