package sign

import (
	"encoding/json"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/signature"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/publicuser"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
)

type Signer struct {
	Id         string   `json:",omitempty"`
	Name       string   `json:",omitempty"`
	Contact    string   `json:",omitempty"`
	Method     string   `json:",omitempty"`
	IsSent     bool     `json:",omitempty"`
	Status     string   `json:",omitempty"`
	DocumentId string   `json:",omitempty"`
	Tags       []string `json:",omitempty"`
	Created    string   `json:",omitempty"`
	Updated    string   `json:",omitempty"`
}

type Document struct {
	Id         string                `json:",omitempty"`
	Content    string                `json:",omitempty"`
	Status     string                `json:",omitempty"`
	Signers    []Signer              `json:",omitempty"`
	Signatures []signature.Signature `json:",omitempty"`
}

var object Document
var resourceDocument = map[string]string{"name": "Document"}

func Get(id string) (Document, Error.StarkErrors) {
	get, err := rest.GetId(
		utils.SdkVersion,
		hosts.Sign,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		publicuser.PublicUser{Environment: "sandbox"},
		resourceDocument,
		id,
		nil,
	)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}
