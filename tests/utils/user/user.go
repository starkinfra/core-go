package user

import (
	"core-go/starkcore/user/organization"
	"core-go/starkcore/user/user"
	"os"
)

var ExampleProject = user.Users{
	os.Getenv("PROJECT_ID"),  // Id: "8888888888888888"
	os.Getenv("PRIVATE_KEY"), // Pem: "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
	"sandbox",                // Environment: "sandbox"
}

var ExampleOrganization = organization.Organization{
	"",
	user.Users{Id: os.Getenv("PROJECT_ID")}, // Id: "8888888888888888"
	user.Users{Pem: os.Getenv("PRIVATE_KEY")},         // Pem: "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
	user.Users{Environment: os.Getenv("ENVIRONMENT")}, // Environment: "sandbox"
}
