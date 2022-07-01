package test

import (
	"core-go/starkcore/utils/rest"
	"core-go/tests/utils/user"
	"fmt"
	"testing"
)

func TestSuccess(t *testing.T) {
	transactions, _ := rest.GetPage(
		"0.0.0",
		"",
		"v2",
		user.ExampleProject,
		"",
		"pt-BR",
		"",
		nil,
	)

	fmt.Sprintf("%v", transactions)
}
