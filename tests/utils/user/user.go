package user

import (
	"github.com/starkinfra/core-go/starkcore/user/organization"
	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
)

var ExampleProjectBank = project.Project{
	Id:          "6253551860842496",                          // Id: "8888888888888888"
	PrivateKey:  checks.CheckPrivateKey("-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIHH8qC9IUc8Tjxm7FN15kFa/4YBqUsOcvnYYoyUVzVesoAcGBSuBBAAK\noUQDQgAEbYawjeUrCDJeb5qYqyY+WTgw+BcToHbxSP6rskxIqxyrUkNbIAWBUWA6\nwZQtSLOrvr8Bfy2LQSLxypO+rRgZGw==\n-----END EC PRIVATE KEY-----"), // "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
	Environment: checks.CheckEnvironment("sandbox"),               // Environment: "sandbox"
}

var ExampleProjectInfra = project.Project{
	Id:          "6253551860842496",                          // Id: "8888888888888888"
	PrivateKey:  checks.CheckPrivateKey("-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIHH8qC9IUc8Tjxm7FN15kFa/4YBqUsOcvnYYoyUVzVesoAcGBSuBBAAK\noUQDQgAEbYawjeUrCDJeb5qYqyY+WTgw+BcToHbxSP6rskxIqxyrUkNbIAWBUWA6\nwZQtSLOrvr8Bfy2LQSLxypO+rRgZGw==\n-----END EC PRIVATE KEY-----"), // "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
	Environment: checks.CheckEnvironment("sandbox"),               // Environment: "sandbox"
}

var ExampleOrganization = organization.Organization{
	Id:          "",                     // Id: "8888888888888888"
	PrivateKey:  "", // "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
	Environment: "",               // Environment: "sandbox"
}
