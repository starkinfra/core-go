package test

import (
	resource2 "core-go/starkcore/utils/resource"
	"core-go/starkcore/utils/rest"
	"core-go/tests/utils/user"
	"fmt"
	"testing"
)

type Transaction struct {
	Amount      int
	Description string
	ExternalId  string
	ReceiverId  string
	Tags        []string
	SenderId    string
	Source      string
	Id          string
	Fee         int
	Balance     int
	Created     int
}

var resource = resource2.Resource{
	Class: Transaction{},
	Name:  "Transaction",
}

func TestSuccess(t *testing.T) {
	transactions, _ := rest.GetPage(
		"0.0.0",
		"",
		"v2",
		user.ExampleProject,
		resource,
		"pt-BR",
		"",
		nil,
	)

	fmt.Sprintf("%v", transactions)
}
