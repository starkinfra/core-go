package invoice

import (
	"encoding/json"
	"fmt"
	"time"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
	User "github.com/starkinfra/core-go/tests/utils/user"
)

type Invoice struct {
	Amount         int                      `json:",omitempty"`
	Name           string                   `json:",omitempty"`
	TaxId          string                   `json:",omitempty"`
	Due            *time.Time               `json:",omitempty"`
	Expiration     int                      `json:",omitempty"`
	Fine           float64                  `json:",omitempty"`
	Interest       float64                  `json:",omitempty"`
	Discounts      []map[string]interface{} `json:",omitempty"`
	Tags           []string                 `json:",omitempty"`
	Descriptions   []map[string]string      `json:",omitempty"`
	Pdf            string                   `json:",omitempty"`
	Link           string                   `json:",omitempty"`
	NominalAmount  int                      `json:",omitempty"`
	FineAmount     int                      `json:",omitempty"`
	InterestAmount int                      `json:",omitempty"`
	DiscountAmount int                      `json:",omitempty"`
	Id             string                   `json:",omitempty"`
	Brcode         string                   `json:",omitempty"`
	Status         string                   `json:",omitempty"`
	Fee            int                      `json:",omitempty"`
	TransactionIds []string                 `json:",omitempty"`
	Created        *time.Time               `json:",omitempty"`
	Updated        string                   `json:",omitempty"`
}

var ResourceInvoice = map[string]string{"name": "Invoice"}

type Payment struct {
	Amount        int    `json:",omitempty"`
	Name          string `json:",omitempty"`
	TaxId         string `json:",omitempty"`
	BankCode      string `json:",omitempty"`
	BranchCode    string `json:",omitempty"`
	AccountNumber string `json:",omitempty"`
	AccountType   string `json:",omitempty"`
	EndToEndId    string `json:",omitempty"`
	Method        string `json:",omitempty"`
}

var subResourcePayment = map[string]string{"name": "Payment"}

func Create(invoices []Invoice) ([]Invoice, Error.StarkErrors) {
	create, err := rest.PostMulti(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		ResourceInvoice,
		invoices,
		nil,
	)
	if err.Errors != nil {
		return []Invoice{}, err
	}
	unmarshalError := json.Unmarshal(create, &invoices)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return invoices, err
}

func CreateWithUser(invoices []Invoice, user user.User) ([]Invoice, Error.StarkErrors) {
	create, err := rest.PostMulti(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		user,
		ResourceInvoice,
		invoices,
		nil,
	)
	if err.Errors != nil {
		return []Invoice{}, err
	}
	unmarshalError := json.Unmarshal(create, &invoices)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return invoices, err
}

func Query(params map[string]interface{}, user user.User) (chan Invoice, chan Error.StarkErrors) {
	var invoice Invoice
	b := make(chan Invoice)
	erroChannel := make(chan Error.StarkErrors)
	c, err := rest.GetStream(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		user,
		ResourceInvoice,
		params,
	)
	go func() {
		for {
			select {
				case errors := <-err:
					erroChannel <- errors
					return

				case value := <-c:
					wereByte, _ := json.Marshal(value)
					err := json.Unmarshal(wereByte, &invoice)
					if err != nil {
						print(err)
					}
					b <- invoice
				}
		}
	}()
	return b, erroChannel
}

func Update(id string) (Invoice, Error.StarkErrors) {
	var invoice Invoice
	var invoicePatch = map[string]interface{}{}
	invoicePatch["amount"] = 1

	updated, err := rest.PatchId(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		ResourceInvoice,
		id,
		invoicePatch,
		nil,
	)
	if err.Errors != nil {
		return Invoice{}, err
	}
	unmarshalError := json.Unmarshal(updated, &invoice)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return invoice, err
}

func Qrcode(id string, query map[string]interface{}) ([]byte, Error.StarkErrors) {
	content, err := rest.GetContent(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		ResourceInvoice,
		id,
		"qrcode",
		query,
	)
	return content, err
}

func Pdf(id string) ([]byte, Error.StarkErrors) {
	content, err := rest.GetContent(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		ResourceInvoice,
		id,
		"pdf",
		nil,
	)
	return content, err
}

func GetPayment(id string) (Payment, Error.StarkErrors) {
	var payment Payment
	get, err := rest.GetSubResource(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		ResourceInvoice,
		id,
		subResourcePayment,
		nil,
	)
	if err.Errors != nil {
		return Payment{}, err
	}
	unmarshalError := json.Unmarshal(get, &payment)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return payment, err
}
