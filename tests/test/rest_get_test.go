package test

import (
	"core-go/starkcore/utils/hosts"
	"core-go/starkcore/utils/resource"
	"core-go/starkcore/utils/rest"
	"core-go/tests/utils/user"
	"fmt"
	"testing"
)

type Boleto struct {
	Amount   int `json:"amount"`
	Resource resource.Resource
}

var boleto Boleto

var resource_g = map[string]string{"name": "Boleto"}

func TestSuccessGetPage(t *testing.T) {
	transactions, _ := rest.GetPage(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_g,
		"pt-BR",
	)

	fmt.Sprintf("%v", transactions)
}

//func TestSuccessGetId(t *testing.T) {
//	transactions, _ := rest.GetId(
//		"0.0.0",
//		"",
//		"v2",
//		user.ExampleOrganization,
//		resource_g,
//		"",
//		"pt-BR",
//	)
//
//	fmt.Sprintf("%v", transactions)
//}
//
//func TestSuccessGetContent(t *testing.T) {
//	transactions, _ := rest.GetContent(
//		"0.0.0",
//		"",
//		"v2",
//		user.ExampleOrganization,
//		resource_g,
//		"",
//		"pt-BR",
//		"",
//	)
//
//	fmt.Sprintf("%v", transactions)
//}
//
//func TestSuccessGetRaw(t *testing.T) {
//	transactions, _ := rest.GetRaw(
//		"0.0.0",
//		"",
//		"v2",
//		user.ExampleOrganization,
//		resource_g,
//		"pt-BR",
//	)
//
//	fmt.Sprintf("%v", transactions)
//}
//
//func TestSuccessGetStream(t *testing.T) {
//	transactions, _ := rest.GetStream(
//		"0.0.0",
//		"",
//		"v2",
//		user.ExampleOrganization,
//		resource_g,
//		"pt-BR",
//	)
//
//	fmt.Sprintf("%v", transactions)
//}
//
//func TestSuccessGetSubResource(t *testing.T) {
//	transactions, _ := rest.GetSubResource(
//		"0.0.0",
//		"",
//		"v2",
//		user.ExampleOrganization,
//		resource_g,
//		"pt-BR",
//	)
//
//	fmt.Sprintf("%v", transactions)
//}
//
//func TestSuccessGetSubResources(t *testing.T) {
//	transactions, _ := rest.GetSubResources(
//		"0.0.0",
//		"",
//		"v2",
//		user.ExampleOrganization,
//		resource_g,
//	)
//
//	fmt.Sprintf("%v", transactions)
//}
