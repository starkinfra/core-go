package api

import (
	"encoding/json"
	"fmt"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	Log "github.com/starkinfra/core-go/tests/utils/boleto/log"
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"os"
	"math/rand"
	"testing"
	"time"
)

func TestSuccessGetStreamBank(t *testing.T) {
	var boleto Boleto.Boleto
	var params = map[string]interface{}{}
	params["limit"] = 500

	streamChannel, errorChannel := rest.GetStream(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoleto,
		params,
	)

	for {
		select {
		case err, ok := <-errorChannel:
			if !ok {
				return
			}
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case data, ok := <-streamChannel:
			if !ok {
				return
			}
			contentByte, _ := json.Marshal(data)
			err := json.Unmarshal(contentByte, &boleto)
			if err != nil {
				t.Errorf("err: %s", err.Error())
			}
			if boleto.Id == "" {
				t.Errorf("boleto.Id is empty")
			}
		}
	}
}

func TestGetStreamWithError(t *testing.T) {
	var params = map[string]interface{}{}
	params["limit"] = 300

	streamChannel, errorChannel := rest.GetStream(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		"pt-PT",
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoleto,
		params,
	)

	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				return
			}
		case <-streamChannel:
			t.Errorf("data is not empty")
		}
	}
}

func TestSuccessGetPage(t *testing.T) {
	var boletos []Boleto.Boleto
	pageData, cursor, err := rest.GetPage(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoleto,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	if cursor == "" {
		t.Errorf("cursor is empty")
	}

	unmarshalError := json.Unmarshal(pageData, &boletos)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
	}

	for _, boleto := range boletos {
		if boleto.Id == "" {
			t.Errorf("boleto.Id is empty")
		}
	}
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
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(get, &boleto)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
	}
	if boleto.Id == "" {
		t.Errorf("boleto.Id is empty")
	}
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
		for _, err := range requestError.Errors {
			t.Errorf("code: %s, message: %s", err.Code, err.Message)
		}
	}
	filename := fmt.Sprintf("%v%v.%v", api.Endpoint(utils.ResourceInvoice), "6543381610102784", "pdf")
	err := os.WriteFile(filename, request, 0666)
	if err != nil {
		t.Errorf("err: %s", err)
	}
}

func TestSuccessGetRaw(t *testing.T) {
	data := map[string]interface{}{}
	pem, err := rest.GetRaw(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		"public-key",
		User.ExampleProjectBank,
		nil,
		"",
		true,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	unmarshalError := json.Unmarshal(pem.Content, &data)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
	}

	if data["publicKeys"] == nil {
		t.Errorf("data['publicKeys'] is nil")
	}
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
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	unmarshalPaymentError := json.Unmarshal(subResource, &payment)
	if unmarshalPaymentError != nil {
		t.Errorf("unmarshalPaymentError: %s", unmarshalPaymentError)
	}
	if payment.Amount == 0 {
		t.Errorf("payment.Amount is 0")
	}
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
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalLogError := json.Unmarshal(get, &log)
	if unmarshalLogError != nil {
		t.Errorf("unmarshalLogError: %s", unmarshalLogError)
	}
	if log.Id == "" {
		t.Errorf("log.Id is empty")
	}
}

func TestSuccessGetStreamBankHolmes(t *testing.T) {
	type BoletoHolmes struct {
		BoletoId string `json:",omitempty"`
		Tags     []string `json:",omitempty"`
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
	logs := make(chan LogHolmes)

	var params = map[string]interface{}{}
	params["after"] = "2022-11-16"
	params["limit"] = 203


	streamChannel, errorChannel := rest.GetStream(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceHolmesLog,
		params,
	)
	go func() {
		defer close(logs)
		for {
			select {
			case err := <-errorChannel:
				if err.Errors != nil {
					for _, e := range err.Errors {
						t.Errorf("code: %s, message: %s", e.Code, e.Message)
					}
				}
			case data, ok := <-streamChannel:
				if !ok {
					return
				}
				contentByte, _ := json.Marshal(data)
				err := json.Unmarshal(contentByte, &log)
				if err != nil {
					t.Errorf("err: %s", err.Error())
				}
				logs <- log
			}
		}
	}()

	for log := range logs {
		if log.Id == "" {
			t.Errorf("log.Id is empty")
		}
	}
}

func TestSuccessGetBalanceStream(t *testing.T) {

	type Balance struct {
		Id       string `json:",omitempty"`
		Amount   int    `json:",omitempty"`
		Currency string `json:",omitempty"`
		Updated  *time.Time
	}

	var balance Balance

	balances := make(chan Balance)

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	streamChannel, errorChannel := rest.GetStream(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBalance,
		params,
	)

	go func() {
		defer close(balances)
		for {
			select {
			case err := <-errorChannel:
				if err.Errors != nil {
					for _, e := range err.Errors {
						t.Errorf("code: %s, message: %s", e.Code, e.Message)
					}
				}
			case data, ok := <-streamChannel:
				if !ok {
					return
				}
				contentByte, _ := json.Marshal(data)
				err := json.Unmarshal(contentByte, &balance)
				if err != nil {
					t.Errorf("err: %s", err.Error())
				}
				balances <- balance
			}
		}
	}()

	for balance := range balances {
		if balance.Id == "" {
			t.Errorf("balance.Id is empty")
		}
	}
}
