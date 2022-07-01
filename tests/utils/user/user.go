package user

import (
	"core-go/starkcore/user/organization"
	"core-go/starkcore/user/project"
	"core-go/starkcore/user/users"
	"os"
)

var ExampleProject = project.Projects{
	"",
	"",
	"sandbox",
	os.Getenv("SANDBOX_PROJECT_ID"),          //"8888888888888888", Pem:
	os.Getenv("SANDBOX_PROJECT_PRIVATE_KEY"), // # "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
	}

}

var ExampleOrganization = organization.Organization{
	"",
	"organization",
	os.Getenv("SANDBOX_PROJECT_ID"),          //"8888888888888888",
	os.Getenv("SANDBOX_PROJECT_PRIVATE_KEY"), // # "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
}
