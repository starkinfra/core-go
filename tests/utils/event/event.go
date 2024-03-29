package event

import (
	"encoding/json"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
	User "github.com/starkinfra/core-go/tests/utils/user"
)

//	Webhook Event struct
//
//	An Event is the notification received from the subscription to the Webhook.
//	Events cannot be created, but may be retrieved from the Stark Bank API to
//	list all generated updates on entities.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the event is created. ex: "5656565656565656"
//	- Log [Log]: A Log struct from one of the subscribed services (TransferLog, InvoiceLog, DepositLog, BoletoLog, BoletoHolmesLog, BrcodePaymentLog, BoletoPaymentLog, UtilityPaymentLog, TaxPaymentLog or DarfPaymentLog)
//	- Created [string]: Creation datetime for the notification event.
//	- IsDelivered [bool]: True if the event has been successfully delivered to the user url. ex: False
//	- Subscription [string]: Service that triggered this event. ex: "transfer", "utility-payment"
//	- WorkspaceId [string]: Id of the Workspace that generated this event. Mostly used when multiple Workspaces have Webhooks registered to the same endpoint. ex: "4545454545454545"

type Event struct {
	Id           string      `json:",omitempty"`
	Log          interface{} `json:",omitempty"`
	Created      string      `json:",omitempty"`
	IsDelivered  bool        `json:",omitempty"`
	Subscription string      `json:",omitempty"`
	WorkspaceId  string      `json:",omitempty"`
}

var object Event
var resourceEvent = map[string]string{"name": "Event"}

func Query(params map[string]interface{}) chan Event {
	b := make(chan Event)

	c := rest.GetStream(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		resourceEvent,
		params,
	)
	go func() {
		for were := range c {
			wereByte, _ := json.Marshal(were)
			err := json.Unmarshal(wereByte, &object)
			if err != nil {
				print(err)
			}
			b <- object
		}
		close(b)
	}()
	return b
}
