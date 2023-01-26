package project

import (
	"fmt"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/privatekey"
)

type Project struct {
	Id          string
	PrivateKey  string
	Environment string
}

//	Project object
//
//	The Project object is an authentication entity for the SDK that is permanently
//	linked to a specific Workspace.
//	All requests to the Stark Bank API must be authenticated via an SDK user,
//	which must have been previously created at the Stark Bank or Stark Infra websites
//	[https://web.sandbox.starkbank.com] or [https://web.starkbank.com]
//	before you can use it in this SDK. Projects may be passed as the user parameter on
//	each request or may be defined as the default user at the start (See README).
//
//	Parameters (required):
//	- Id [string]: unique id required to identify project. ex: "5656565656565656"
//	- PrivateKey [string]: PEM string of the private key linked to the project. ex: "-----BEGIN PUBLIC KEY-----\nMFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEyTIHK6jYuik6ktM9FIF3yCEYzpLjO5X/\ntqDioGM+R2RyW0QEo+1DG8BrUf4UXHSvCjtQ0yLppygz23z0yPZYfw==\n-----END PUBLIC KEY-----"
//	- Environment [string]: environment where the project is being used. ex: "sandbox" or "production"
//
//	Attributes (return-only):
//	- Name [string]: project name. ex: "MyProject"
//	- AllowedIps [slice of strings]: list containing the strings of the ips allowed to make requests on behalf of this project. ex: ["190.190.0.50"]
//	- Pem [string]: private key in pem format. ex: "-----BEGIN PUBLIC KEY-----\nMFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEyTIHK6jYuik6ktM9FIF3yCEYzpLjO5X/\ntqDioGM+R2RyW0QEo+1DG8BrUf4UXHSvCjtQ0yLppygz23z0yPZYfw==\n-----END PUBLIC KEY-----"

func (p Project) GetAcessId() string {
	return fmt.Sprintf("project/%v", p.Id)
}

func (p Project) GetEnvironment() string {
	return fmt.Sprintf("%v", p.Environment)
}

func (p Project) GetPrivateKey() *privatekey.PrivateKey {
	result := privatekey.FromPem(p.PrivateKey)
	return &result
}
