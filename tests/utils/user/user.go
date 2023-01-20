package user

import (
	"github.com/starkinfra/core-go/starkcore/user/organization"
	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
	"os"
)

var ResourceBoletoLog = map[string]string{"name": "BoletoLog"}
var ResourceBoleto = map[string]string{"name": "Boleto"}
var ResourceUtilityPayment = map[string]string{"name": "UtilityPayment"}
var SubResourcePayment = map[string]string{"name": "Payment"}
var ResourcePixDomain = map[string]string{"name": "PixDomain"}
var ResourceCardMethod = map[string]string{"name": "CardMethod"}
var ResourceInvoice = map[string]string{"name": "invoice"}

var Language = "en-US"
var Timeout = 15
var SdkVersion = "0.0.0"
var ApiVersion = "v2"

var ExampleProjectBank = project.Project{
	Id:          os.Getenv("PROJECT_ID"),                           // Id: "8888888888888888"
	PrivateKey:  checks.CheckPrivateKey(os.Getenv("PRIVATE_KEY")),  // "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
	Environment: checks.CheckEnvironment(os.Getenv("ENVIRONMENT")), // Environment: "sandbox"
}

var ExampleProjectInfra = project.Project{
	Id:          os.Getenv("PROJECT_ID"),                           // Id: "8888888888888888"
	PrivateKey:  checks.CheckPrivateKey(os.Getenv("PRIVATE_KEY")),  // "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
	Environment: checks.CheckEnvironment(os.Getenv("ENVIRONMENT")), // Environment: "sandbox"
}

var ExampleOrganization = organization.Organization{
	Id:          os.Getenv("ORGANIZATION_ID"),                      // Id: "8888888888888888"
	PrivateKey:  checks.CheckPrivateKey(os.Getenv("PRIVATE_KEY")),  // "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
	Environment: checks.CheckEnvironment(os.Getenv("ENVIRONMENT")), // Environment: "sandbox"
}
