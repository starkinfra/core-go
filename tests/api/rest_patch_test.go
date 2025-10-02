package api

import (
	"encoding/json"
	Invoice "github.com/starkinfra/core-go/tests/utils/invoice"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
	"github.com/starkinfra/core-go/tests/utils/examples"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"testing"
)

func TestSuccessPatch(t *testing.T) {
	var invoice Invoice.Invoice
	var patchData = map[string]interface{}{}
	var amount = 200
	patchData["amount"] = amount

	createdInvoice, err := Invoice.Create(examples.ExampleInvoice())
	if err.Errors != nil {
		t.Errorf("err: %s", err.Errors)
	}

	patch, err := rest.PatchId(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceInvoice,
		createdInvoice[0].Id,
		patchData,
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	unmarshalError := json.Unmarshal(patch, &invoice)
	if unmarshalError != nil {
		t.Errorf("unmarshalError: %s", unmarshalError)
	}

	if invoice.Amount != amount {
		t.Errorf("invoice.Amount is not %d", amount)
	}
}

func TestFailedPatch(t *testing.T) {
	var patchData = map[string]interface{}{}
	patchData["amount"] = 8

	_, err := rest.PatchId(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceInvoice,
		"5806698551836672",
		patchData,
		nil,
	)
	if err.Errors != nil {
		return
	}
	t.Errorf("unmarshalError: %s", err.Errors)
}
