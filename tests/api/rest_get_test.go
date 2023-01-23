package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"

	"github.com/starkinfra/core-go/starkcore/utils/api"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	Log "github.com/starkinfra/core-go/tests/utils/boleto/log"
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
	User "github.com/starkinfra/core-go/tests/utils/user"
)

func TestSuccessGetStreamBank(t *testing.T) {
	var boletos []Boleto.Boleto
	data, err := rest.GetStream(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceBoleto,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	queryError := json.Unmarshal(data, &boletos)
	if queryError != nil {
		fmt.Println(queryError)
	}
	for _, boleto := range boletos {
		fmt.Println(boleto.Id)
	}
}

func TestSuccessGetPage(t *testing.T) {
	payments, cursor, err := rest.GetPage(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceUtilityPayment,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println(cursor)
	fmt.Println(string(payments))
}

func TestSuccessGetId(t *testing.T) {
	var boleto Boleto.Boleto
	var boletos []Boleto.Boleto

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	query, err := rest.GetStream(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceBoleto,
		params,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	queryError := json.Unmarshal(query, &boletos)
	if queryError != nil {
		fmt.Println(queryError)
	}

	get, err := rest.GetId(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceBoleto,
		boletos[rand.Intn(params["limit"].(int))].Id,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(get, &boleto)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	fmt.Println(boleto.Id)
}

func TestSuccessGetContent(t *testing.T) {

	var invoices []Invoice.Invoice

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	query, queryError := rest.GetStream(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceInvoice,
		params,
	)
	if queryError.Errors != nil {
		for _, e := range queryError.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(query, &invoices)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}

	request, requestError := rest.GetContent(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceInvoice,
		invoices[rand.Intn(params["limit"].(int))].Id,
		"pdf",
		nil,
	)
	if requestError.Errors != nil {
		for _, erro := range requestError.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	filename := fmt.Sprintf("%v%v.%v", api.Endpoint(User.ResourceInvoice), invoices[rand.Intn(params["limit"].(int))].Id, "pdf")
	err := ioutil.WriteFile(filename, request, 0666)
	if err != nil {
		fmt.Println("errrrrrrr", err)
	}
}

func TestSuccessGetRaw(t *testing.T) {
	pem, err := rest.GetRaw(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		"public-key",
		User.ExampleProjectBank,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println(pem)
}

func TestSuccessGetStreamInfra(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 3

	data, err := rest.GetStream(
		User.SdkVersion,
		hosts.Infra,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectInfra,
		User.ResourcePixDomain,
		params,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println(string(data))
}

func TestSuccessGetStreamInfraCardMethod(t *testing.T) {
	data, err := rest.GetStream(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectInfra,
		User.ResourceCardMethod,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println(string(data))
}

func TestSuccessGetSubResource(t *testing.T) {
	var invoices []Invoice.Invoice
	var payment Invoice.Payment

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	query, queryError := rest.GetStream(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceInvoice,
		params,
	)
	if queryError.Errors != nil {
		for _, e := range queryError.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(query, &invoices)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}

	subResource, err := rest.GetSubResource(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceInvoice,
		invoices[rand.Intn(params["limit"].(int))].Id,
		User.SubResourcePayment,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	unmarshalPaymentError := json.Unmarshal(subResource, &payment)
	if unmarshalPaymentError != nil {
		fmt.Println(unmarshalPaymentError)
	}
	fmt.Println(payment.Amount)
}

func TestSuccessGetLogId(t *testing.T) {
	var logs []Log.Log
	var log Log.Log

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	query, queryError := rest.GetStream(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceBoletoLog,
		params,
	)
	if queryError.Errors != nil {
		for _, e := range queryError.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(query, &logs)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}

	get, err := rest.GetId(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceBoletoLog,
		logs[rand.Intn(params["limit"].(int))].Id,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalLogError := json.Unmarshal(get, &log)
	if unmarshalLogError != nil {
		fmt.Println(unmarshalLogError)
	}
	fmt.Println(log.Id)
}

func TestSuccessGetHolmesLogId(t *testing.T) {
	id, err := rest.GetId(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		Log.ResourceHolmesLog,
		"5682651440611328",
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println(string(id))
}

func TestSuccessGetStreamBankHolmes(t *testing.T) {
	type BoletoHolmes struct {
		BoletoId string `json:",omitempty"`
		Tags     string `json:",omitempty"`
		Id       string `json:",omitempty"`
		Status   string `json:",omitempty"`
		Result   string `json:",omitempty"`
		Created  *time.Time
		Updated  *time.Time
	}

	type LogHolmes struct {
		Id      string       `json:",omitempty"`
		Holmes  BoletoHolmes `json:",omitempty"`
		Errors  []string     `json:",omitempty"`
		Type    string       `json:",omitempty"`
		Created *time.Time   `json:",omitempty"`
	}

	var logs []LogHolmes

	var params = map[string]interface{}{}
	params["after"] = "2022-11-16"
	params["limit"] = 203

	get, err := rest.GetStream(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		Log.ResourceHolmesLog,
		params,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(get, &logs)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	for _, log := range logs {
		fmt.Println(log.Id)
	}
}

func TestBalanceStream(t *testing.T) {

	var resourceBalance = map[string]string{"name": "Balance"}

	type Balance struct {
		Id       string `json:",omitempty"`
		Amount   int    `json:",omitempty"`
		Currency string `json:",omitempty"`
		Updated  *time.Time
	}

	var balances []Balance

	get, err := rest.GetStream(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		resourceBalance,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(get, &balances)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	for _, balance := range balances {
		fmt.Println(balance)
	}
}
