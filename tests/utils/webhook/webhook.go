package webhook

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
	User "github.com/starkinfra/core-go/tests/utils/user"
)

type Webhook struct {
	Url           string   `json:",omitempty"`
	Subscriptions []string `json:",omitempty"`
	Id            string   `json:",omitempty"`
}

var resourceWebhook = map[string]string{"name": "Webhook"}

func Create(webhook Webhook) (Webhook, Error.StarkErrors) {
	create, err := rest.PostSingle(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		resourceWebhook,
		webhook,
		nil,
	)
	if err.Errors != nil {
		return Webhook{}, err
	}
	unmarshalError := json.Unmarshal(create, &webhook)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return webhook, err
}
