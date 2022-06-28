package user

import (
	"core-go/starkcore/user/organization"
	"core-go/starkcore/user/project"
	"os"
)

var exampleProject = project.Project{
	Environment: "sandbox",
	Id:          os.Getenv("SANDBOX_PROJECT_ID"),          //"8888888888888888",
	PrivateKey:  os.Getenv("SANDBOX_PROJECT_PRIVATE_KEY"), // # "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
}

var exampleOrganization = organization.Organization{
	Environment: "organization",
	Id:          os.Getenv("SANDBOX_PROJECT_ID"),          //"8888888888888888",
	PrivateKey:  os.Getenv("SANDBOX_PROJECT_PRIVATE_KEY"), // # "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
}
