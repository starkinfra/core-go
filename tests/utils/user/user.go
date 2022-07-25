package user

import (
	"core-go/starkcore/user/project"
	"core-go/starkcore/user/user"
	"core-go/starkcore/utils/checks"
	"os"
)

var ExampleProject = project.Projects{
	"",                                      // Name: "Apollo 13"
	nil,                                     // AllowedIps: ["128.000.01", "127.000.01"]
	user.Users{Id: os.Getenv("PROJECT_ID")}, //os.Getenv("PROJECT_ID"), // Id: "8888888888888888"
	user.Users{Pem: checks.CheckPrivateKey(os.Getenv("PRIVATE_KEY"))},          //os.Getenv("PRIVATE_KEY"), // Pem: "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"} //os.Getenv("PRIVATE_KEY")},         // Pem: "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
	user.Users{Environment: checks.CheckEnvironment(os.Getenv("ENVIRONMENT"))}, //os.Getenv("ENVIRONMENT")}, // Environment: "sandbox"
}

//var ExampleOrganization = organization.Organization{
//	"",
//	user.Users{Id: os.Getenv("PROJECT_ID")}, // Id: "8888888888888888"
//	user.Users{Pem: os.Getenv("PRIVATE_KEY")}, // Pem: "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"
//	user.Users{Environment: os.Getenv("ENVIRONMENT")}, // Environment: "sandbox"
//}
