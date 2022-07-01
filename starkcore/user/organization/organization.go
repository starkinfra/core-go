package organization

import (
	"core-go/starkcore/user/users"
	"fmt"
)

type Organization struct {
	WorkspaceId string
	Id          users.User
	PrivateKey  users.User
	Environment users.User
}

//  Organization object
//  The Organization object is an authentication entity for the SDK that
//  represents your entire Organization, being able to access any Workspace
//  underneath it and even create new Workspaces. Only a legal representative
//  of your organization can register or change the Organization credentials.
//  All requests to the Stark Bank and Stark Infra API must be authenticated via an SDK users,
//  which must have been previously created at the Stark Bank or Stark Infra websites
//  [https://web.sandbox.starkbank.com] or [https://web.starkbank.com]
//  before you can use it in this SDK. Organizations may be passed as the users parameter on
//  each request or may be defined as the default users at the start (See README).
//  If you are accessing a specific Workspace using Organization credentials, you should
//  specify the workspace ID when building the Organization object or by request, using
//  the Organization.replace(organization, workspace_id) function, which creates a copy of the organization
//  object with the altered workspace ID. If you are listing or creating new Workspaces, the
//  workspace_id should be None.
//
//  Parameters (required):
//  - id [string]: unique id required to identify organization. ex: "5656565656565656"
//  - private_key [EllipticCurve.Organization()]: PEM string of the private key linked to the organization. ex: "-----BEGIN PUBLIC KEY-----\nMFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEyTIHK6jYuik6ktM9FIF3yCEYzpLjO5X/\ntqDioGM+R2RyW0QEo+1DG8BrUf4UXHSvCjtQ0yLppygz23z0yPZYfw==\n-----END PUBLIC KEY-----"
//  - environment [string]: environment where the organization is being used. ex: "sandbox" or "production"
//  - workspace_id [string]: unique id of the accessed Workspace, if any. ex: None or "4848484848484848"
//
//  Attributes (return-only):
//  - pem [string]: private key in pem format. ex: "-----BEGIN PUBLIC KEY-----\nMFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEyTIHK6jYuik6ktM9FIF3yCEYzpLjO5X/\ntqDioGM+R2RyW0QEo+1DG8BrUf4UXHSvCjtQ0yLppygz23z0yPZYfw==\n-----END PUBLIC KEY-----"

func (o Organization) AccessId() string {
	if o.WorkspaceId == "" {
		return fmt.Sprintf("organization/%i", o.Id)
	}
	return fmt.Sprintf("organization/%i/workspace/%w", o.Id, o.WorkspaceId)
}

func Replace(workspaceId string, organization Organization) Organization {
	return Organization{
		WorkspaceId: workspaceId,
		Id:          organization.Id,
		PrivateKey:  organization.PrivateKey,
		Environment: organization.Environment,
	}
}
