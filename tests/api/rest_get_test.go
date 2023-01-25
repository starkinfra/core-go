package api

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	Log "github.com/starkinfra/core-go/tests/utils/boleto/log"
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"
)

func TestSuccessGetStreamBank(t *testing.T) {
	var boleto Boleto.Boleto
	b := make(chan Boleto.Boleto)
	c := make(chan map[string]interface{})
	e := make(chan Error.StarkError)
	f := make(chan Error.StarkError)

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	go rest.GetStream(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoleto,
		params,
		c,
		e,
	)
	if e != nil {
		go func() {
			for were := range e {
				example := were
				f <- example
			}
			close(f)
		}()
	}
	go func() {
		for were := range c {
			wereByte, _ := json.Marshal(were)
			err := json.Unmarshal(wereByte, &boleto)
			if err != nil {
				print(err)
			}
			b <- boleto
		}
		close(b)
	}()

	for entity := range b {
		fmt.Println(entity.Id)
	}
}

func TestSuccessGetPage(t *testing.T) {
	payments, cursor, err := rest.GetPage(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceUtilityPayment,
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

	get, err := rest.GetId(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoleto,
		"4537841761648640",
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
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	request, requestError := rest.GetContent(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceInvoice,
		"6543381610102784",
		"pdf",
		nil,
	)
	if requestError.Errors != nil {
		for _, erro := range requestError.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	filename := fmt.Sprintf("%v%v.%v", api.Endpoint(utils.ResourceInvoice), "6543381610102784", "pdf")
	err := ioutil.WriteFile(filename, request, 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func TestSuccessGetRaw(t *testing.T) {
	pem, err := rest.GetRaw(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
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

func TestSuccessGetSubResource(t *testing.T) {
	var payment Invoice.Payment

	subResource, err := rest.GetSubResource(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceInvoice,
		"6543381610102784",
		utils.SubResourcePayment,
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
	var log Log.Log

	get, err := rest.GetId(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoletoLog,
		"4919717689032704",
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
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceHolmesLog,
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

	var log LogHolmes

	var params = map[string]interface{}{}
	params["after"] = "2022-11-16"
	params["limit"] = 203

	b := make(chan LogHolmes)
	c := make(chan map[string]interface{})
	e := make(chan Error.StarkError)
	f := make(chan Error.StarkError)
	go rest.GetStream(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceHolmesLog,
		params,
		c,
		e,
	)
	if e != nil {
		go func() {
			for were := range e {
				example := were
				f <- example
			}
			close(f)
		}()
	}
	go func() {
		for were := range c {
			wereByte, _ := json.Marshal(were)
			err := json.Unmarshal(wereByte, &log)
			if err != nil {
				print(err)
			}
			b <- log
		}
		close(b)
	}()

	for entity := range b {
		fmt.Println(entity.Id)
	}
}

func TestBalanceStream(t *testing.T) {

	type Balance struct {
		Id       string `json:",omitempty"`
		Amount   int    `json:",omitempty"`
		Currency string `json:",omitempty"`
		Updated  *time.Time
	}

	var balance Balance

	b := make(chan Balance)
	c := make(chan map[string]interface{})
	e := make(chan Error.StarkError)
	f := make(chan Error.StarkError)

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	go rest.GetStream(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBalance,
		params,
		c,
		e,
	)
	if e != nil {
		go func() {
			for were := range e {
				example := were
				f <- example
			}
			close(f)
		}()
	}
	go func() {
		for were := range c {
			wereByte, _ := json.Marshal(were)
			err := json.Unmarshal(wereByte, &balance)
			if err != nil {
				print(err)
			}
			b <- balance
		}
		close(b)
	}()

	for entity := range b {
		fmt.Println(entity.Id)
	}
}
