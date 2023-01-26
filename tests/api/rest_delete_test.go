package api

import (
	"encoding/json"
	"fmt"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"math/rand"
	"testing"
)

func TestSuccessDel(t *testing.T) {
	var boleto Boleto.Boleto

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	deleted, err := rest.DeleteId(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoleto,
		"6678130383126528",
		nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(deleted, &boleto)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}

	fmt.Println("id: ", boleto.Id)

}
