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

var body_pa = struct {
}{}

var resource_p = resource2.Resource{
	Class: Transaction{},
	Name:  "Transaction",
}

func TestSuccessPatch(t *testing.T) {
	transactions, _ := rest.PatchId(
		"0.0.0",
		"",
		"v2",
		user.ExampleOrganization,
		resource_p,
		"",
		body_pa,
		"pt-BR",
	)

	fmt.Sprintf("%v", transactions)
}
