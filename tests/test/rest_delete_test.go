package test

import (
	"core-go/starkcore/utils/hosts"
	"core-go/starkcore/utils/rest"
	"core-go/tests/utils/user"
	"testing"
)

var resource_d = map[string]string{"name": "boleto"}

func TestSuccessDel(t *testing.T) {
	rest.DeleteId(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_d,
		"",
		"",
		"pt-BR",
		15,
	)
}
