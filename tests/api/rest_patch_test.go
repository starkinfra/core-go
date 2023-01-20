package api

import (
	"fmt"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"testing"
)

func TestSuccessPatch(t *testing.T) {
	var patchData = map[string]interface{}{}
	patchData["amount"] = 10

	invoice, err := rest.PatchId(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		User.ResourceInvoice,
		"4765711691939840",
		patchData,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println(string(invoice))
}
