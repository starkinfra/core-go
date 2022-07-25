package organization

import (
	"core-go/starkcore/user/user"
	"fmt"
	"github.com/starkbank/ecdsa-go/ellipticcurve/privatekey"
)

type Organization struct {
	WorkspaceId string
	Id          user.Users
	PrivateKey  user.Users
	Environment user.Users
}

//  Organization object
//  The Organization object is an authentication entity for the SDK that
//  represents your entire Organization, being able to access any Workspace
//  underneath it and even create new Workspaces. Only a legal representative
//  of your organization can register or change the Organization credentials.
//  All requests to the Stark Bank and Stark Infra API must be authenticated via an SDK user,
//  which must have been previously created at the Stark Bank or Stark Infra websites
//  [https://web.sandbox.starkbank.com] or [https://web.starkbank.com]
//  before you can use it in this SDK. Organizations may be passed as the user parameter on
//  each request or may be defined as the default user at the start (See README).
//  If you are accessing a specific Workspace using Organization credentials, you should
//  specify the workspace ID when building the Organization object or by request, using
//  the organization.Replace(workspaceId, organization) function, which creates a copy of the organization
//  object with the altered workspace ID. If you are listing or creating new Workspaces, the
//  workspaceId should be None.
//
//  Parameters (required):
//  - Id [string]: unique id required to identify organization. ex: "5656565656565656"
//  - PrivateKey [string]: PEM string of the private key linked to the organization. ex: "-----BEGIN PUBLIC KEY-----\nMFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEyTIHK6jYuik6ktM9FIF3yCEYzpLjO5X/\ntqDioGM+R2RyW0QEo+1DG8BrUf4UXHSvCjtQ0yLppygz23z0yPZYfw==\n-----END PUBLIC KEY-----"
//  - Environment [string]: environment where the organization is being used. ex: "sandbox" or "production"
//  - WorkspaceId [string]: unique id of the accessed Workspace, if any. ex: None or "4848484848484848"
//
//  Attributes (return-only):
//  - pem [string]: private key in pem format. ex: "-----BEGIN PUBLIC KEY-----\nMFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEyTIHK6jYuik6ktM9FIF3yCEYzpLjO5X/\ntqDioGM+R2RyW0QEo+1DG8BrUf4UXHSvCjtQ0yLppygz23z0yPZYfw==\n-----END PUBLIC KEY-----"

func (o Organization) AccessId() string {
	if o.WorkspaceId == "" {
		return fmt.Sprintf("organization/%v", o.Id.Id)
	}
	return fmt.Sprintf("organization/%v/workspace/%v", o.Id.Id, o.WorkspaceId)
}

func (o Organization) Environments() string {
	return fmt.Sprintf("%v", o.Environment.Environment)
}

func (o Organization) PrivateKeys() *privatekey.PrivateKey {
	result := privatekey.FromPem(o.PrivateKey.Pem)
	return &result
}

func Replace(workspaceId string, organization Organization) Organization {
	return Organization{
		WorkspaceId: workspaceId,
		Id:          organization.Id,
		PrivateKey:  organization.PrivateKey,
		Environment: organization.Environment,
	}
}
