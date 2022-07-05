package test

import (
	resource2 "core-go/starkcore/utils/resource"
	"core-go/starkcore/utils/rest"
	"core-go/tests/utils/user"
	"fmt"
	"testing"
)

type Transaction struct {
	Amount      int      `json:"amount"`
	Description string   `json:"amount"`
	ExternalId  string   `json:"amount"`
	ReceiverId  string   `json:"amount"`
	Tags        []string `json:"amount"`
	SenderId    string   `json:"amount"`
	Source      string   `json:"amount"`
	Id          string   `json:"amount"`
	Fee         int      `json:"amount"`
	Balance     int      `json:"amount"`
	Created     int      `json:"amount"`
}

var resource_g = resource2.Resource{
	Class: Transaction{},
	Name:  "Transaction",
}

func TestSuccessGetPage(t *testing.T) {
	transactions, _ := rest.GetPage(
		"0.0.0",
		"",
		"v2",
		user.ExampleOrganization,
		resource_g,
		"pt-BR",
	)

	fmt.Sprintf("%v", transactions)
}

func TestSuccessGetId(t *testing.T) {
	transactions, _ := rest.GetId(
		"0.0.0",
		"",
		"v2",
		user.ExampleOrganization,
		resource_g,
		"",
		"pt-BR",
	)

	fmt.Sprintf("%v", transactions)
}

func TestSuccessGetContent(t *testing.T) {
	transactions, _ := rest.GetContent(
		"0.0.0",
		"",
		"v2",
		user.ExampleOrganization,
		resource_g,
		"",
		"pt-BR",
		"",
	)

	fmt.Sprintf("%v", transactions)
}

func TestSuccessGetRaw(t *testing.T) {
	transactions, _ := rest.GetRaw(
		"0.0.0",
		"",
		"v2",
		user.ExampleOrganization,
		resource_g,
		"pt-BR",
	)

	fmt.Sprintf("%v", transactions)
}

func TestSuccessGetStream(t *testing.T) {
	transactions, _ := rest.GetStream(
		"0.0.0",
		"",
		"v2",
		user.ExampleOrganization,
		resource_g,
		"pt-BR",
	)

	fmt.Sprintf("%v", transactions)
}

func TestSuccessGetSubResource(t *testing.T) {
	transactions, _ := rest.GetSubResource(
		"0.0.0",
		"",
		"v2",
		user.ExampleOrganization,
		resource_g,
		"pt-BR",
	)

	fmt.Sprintf("%v", transactions)
}

func TestSuccessGetSubResources(t *testing.T) {
	transactions, _ := rest.GetSubResources(
		"0.0.0",
		"",
		"v2",
		user.ExampleOrganization,
		resource_g,
	)

	fmt.Sprintf("%v", transactions)
}
