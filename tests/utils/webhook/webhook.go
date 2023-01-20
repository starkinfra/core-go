package webhook

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
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
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		resourceWebhook,
		webhook,
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
