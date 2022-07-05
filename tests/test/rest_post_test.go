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

var resource_po = resource2.Resource{
	Class: Transaction{},
	Name:  "Transaction",
}

var body_po = struct {
}{}

func TestSuccessSingle(t *testing.T) {
	transactions, _ := rest.PostSingle(
		"0.0.0",
		"",
		"v2",
		user.ExampleOrganization,
		resource_po,
		body,
		"",
	)

	fmt.Sprintf("%v", transactions)
}

func TestSuccessMulti(t *testing.T) {
	transactions, _ := rest.PostMulti(
		"0.0.0",
		"",
		"v2",
		user.ExampleOrganization,
		resource_po,
		body_po,
		"",
	)

	fmt.Sprintf("%v", transactions)
}
