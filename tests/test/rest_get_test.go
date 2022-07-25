package test

import (
	"core-go/starkcore/utils/hosts"
	"core-go/starkcore/utils/rest"
	"core-go/tests/utils/user"
	"testing"
)

type Recursos struct {
	Class Boleto
	Name  string
}

var resource = Recursos{Class: Boleto{}, Name: "Boleto"}

func TestSuccessGetPage(t *testing.T) {
	rest.GetPage(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource,
		"pt-BR",
		15,
		nil,
	)
}

func TestSuccessGetId(t *testing.T) {
	rest.GetId(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_g,
		"",
		"pt-BR",
		15,
		nil,
	)
}

func TestSuccessGetContent(t *testing.T) {
	rest.GetContent(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_g,
		"",
		"pt-BR",
		"",
		15,
		nil,
	)
}

func TestSuccessGetRaw(t *testing.T) {
	rest.GetRaw(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		"v2",
		user.ExampleProject,
		"pt-BR",
		15,
		nil,
	)
}

func TestSuccessGetStream(t *testing.T) {
	rest.GetStream(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_g,
		"pt-BR",
		15,
		0,
		nil,
	)
}

func TestSuccessGetSubResource(t *testing.T) {
	rest.GetSubResource(
		"0.0.0",
		hosts.Service.Bank,
		"v2",
		user.ExampleProject,
		resource_g,
		"",
		"pt-BR",
		"",
		15,
		nil,
	)
}
