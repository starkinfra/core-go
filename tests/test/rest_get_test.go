package test

import (
	"core-go/starkcore/utils/hosts"
	"core-go/starkcore/utils/rest"
	"core-go/tests/utils/user"
	"testing"
)

type Boletos struct {
	Boletos []Boleto
}

type Boleto struct {
	Amount        *int
	Name          *string
	TaxId         *string
	StreetLine1   *string
	StreetLine2   *string
	District      *string
	City          *string
	StateCode     *string
	ZipCode       *string
	Due           *string
	Fine          *float32
	Interest      *float32
	OverdueLimit  *int
	Descriptions  *map[string]string
	Discounts     *map[string]string
	Tags          *[]string
	ReceiverName  *string
	ReceiverTaxId *string
	Id            *string
	Fee           *int
	Line          *string
	BarCode       *string
	Transactions  *[]string
	Created       *string
	OurNumber     *string
}

type ResourceR struct {
	Class interface{}
	Name  string
}

var resource = ResourceR{
	Boletos{}, "boletos",
}

func (r ResourceR) GetObject() interface{} {
	return r.Class
}

func (r ResourceR) GetName() string {
	return r.Name
}

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

//func TestSuccessGetId(t *testing.T) {
//	rest.GetId(
//		"0.0.0",
//		hosts.Service.Bank,
//		"v2",
//		user.ExampleProject,
//		resource,
//		"5756110329872384",
//		"pt-BR",
//		15,
//		nil,
//	)
//}

//
//func TestSuccessGetContent(t *testing.T) {
//	rest.GetContent(
//		"0.0.0",
//		hosts.Service.Bank,
//		"v2",
//		user.ExampleProject,
//		resource_g,
//		"",
//		"pt-BR",
//		"",
//		15,
//		nil,
//	)
//}
//
//func TestSuccessGetRaw(t *testing.T) {
//	rest.GetRaw(
//		"0.0.0",
//		hosts.Service.Bank,
//		"v2",
//		"v2",
//		user.ExampleProject,
//		"pt-BR",
//		15,
//		nil,
//	)
//}
//
//func TestSuccessGetStream(t *testing.T) {
//	rest.GetStream(
//		"0.0.0",
//		hosts.Service.Bank,
//		"v2",
//		user.ExampleProject,
//		resource_g,
//		"pt-BR",
//		15,
//		0,
//		nil,
//	)
//}
//
//func TestSuccessGetSubResource(t *testing.T) {
//	rest.GetSubResource(
//		"0.0.0",
//		hosts.Service.Bank,
//		"v2",
//		user.ExampleProject,
//		resource_g,
//		"",
//		"pt-BR",
//		"",
//		15,
//		nil,
//	)
//}
