package product

import (
	"encoding/json"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	Utils "github.com/starkinfra/core-go/tests/utils"
	User "github.com/starkinfra/core-go/tests/utils/user"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"time"
)

//	IssuingProduct struct
//
//	The IssuingProduct struct displays information of registered card products to your Workspace.
//  They represent a group of cards that begin with the same numbers (id) and offer the same product to end customers.
//
//	Attributes (return-only):
//	- Id [string]: Unique card product number (BIN) registered within the card network. ex: "53810200"
//  - Network [string]: Card network flag. ex: "mastercard"
//  - FundingType [string]: Type of funding used for payment. ex: "credit", "debit"
//  - HolderType [string]: Holder type. ex: "business", "individual"
//  - Code [string]: Internal code from card flag informing the product. ex: "MRW", "MCO", "MWB", "MCS"
//  - Created [time.Time]: Creation datetime for the IssuingProduct. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingProduct struct {
	Id          string     `json:",omitempty"`
	Network     string     `json:",omitempty"`
	FundingType string     `json:",omitempty"`
	HolderType  string     `json:",omitempty"`
	Code        string     `json:",omitempty"`
	Created     *time.Time `json:",omitempty"`
}

var resourceIssuingProduct = map[string]string{"name": "IssuingProduct"}

func Query(params map[string]interface{}) (chan IssuingProduct, chan Error.StarkErrors) {
	//	Retrieve IssuingProduct structs
	//
	//	Receive a generator of IssuingProduct structs previously registered in the Stark Infra API
	//
	//	Parameters (required):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.user was set before function call
	//
	//	Parameters (optional):
	//	- limit [int, default 0]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//
	//	Return:
	//	- generator of IssuingBin structs with updated attributes
	var object IssuingProduct
	b := make(chan IssuingProduct)
	erroChannel := make(chan Error.StarkErrors)
	c, err := rest.GetStream(
		Utils.SdkVersion,
		hosts.Infra,
		Utils.ApiVersion,
		Utils.Language,
		Utils.Timeout,
		User.ExampleProjectInfra,
		resourceIssuingProduct,
		params,
	)
	go func(){
		for {
			select{
				case errors := <- err:
					erroChannel <- errors
					return 
	
				case value := <- c:
					
					wereByte, _ := json.Marshal(value)
					err := json.Unmarshal(wereByte, &object)
					if err != nil {
						print(err)
					}
					b <- object
			}
		}
	}()
	return b, erroChannel
}
