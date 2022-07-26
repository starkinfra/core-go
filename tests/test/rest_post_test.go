package test

//
//import (
//	"core-go/starkcore/utils/api"
//	"core-go/starkcore/utils/hosts"
//	"core-go/starkcore/utils/rest"
//	"core-go/tests/utils/user"
//	"testing"
//)
//
//type Boletos struct {
//	Boletos []Boleto
//}
//
//type Boleto struct {
//	Amount        *int
//	Name          *string
//	TaxId         *string
//	StreetLine1   *string
//	StreetLine2   *string
//	District      *string
//	City          *string
//	StateCode     *string
//	ZipCode       *string
//	Due           *string
//	Fine          *float32
//	Interest      *float32
//	OverdueLimit  *int
//	Descriptions  *map[string]string
//	Discounts     *map[string]string
//	Tags          *[]string
//	ReceiverName  *string
//	ReceiverTaxId *string
//	Id            *string
//	Fee           *int
//	Line          *string
//	BarCode       *string
//	Transactions  *[]string
//	Created       *string
//	OurNumber     *string
//}
//
//func TestSuccessSingle(t *testing.T) {
//
//	amount := 1111111111
//	name := "PRIMEIRO TESTE DE POST SINGLE NA TERÇA FEIRA"
//	taxId := "38.446.231/0001-04"
//	streetLine1 := "Kubasch Street, 900"
//	streetLine2 := ""
//	district := "Ronny"
//	city := "Emmet City"
//	stateCode := "SP"
//	zipCode := "01420-020"
//
//	boletosingle := Boleto{
//		Amount:      &amount,
//		Name:        &name,
//		TaxId:       &taxId,
//		StreetLine1: &streetLine1,
//		StreetLine2: &streetLine2,
//		District:    &district,
//		City:        &city,
//		StateCode:   &stateCode,
//		ZipCode:     &zipCode,
//	}
//
//	rest.PostSingle(
//		"0.0.0",
//		hosts.Service.Bank,
//		"v2",
//		user.ExampleProject,
//		resource_g,
//		api.ApiJson(boletosingle),
//		"pt-BR",
//		15,
//		nil,
//	)
//}
//
//func TestSuccessPostMulti(t *testing.T) {
//
//	amount2 := 22222222222
//	name2 := "PRIMEIRO TESTE DE POSTMULTI NA QUINTA FEIRA"
//	taxId2 := "38.446.231/0001-04"
//	streetLine12 := "Kubasch Street, 900"
//	streetLine22 := ""
//	district2 := "Ronny"
//	city2 := "Emmet City"
//	stateCode2 := "SP"
//	zipCode2 := "01420-020"
//
//	amount3 := 33333333333
//	name3 := "SEGUNDO TESTE DE POSTMULTI NA QUINTA FEIRA"
//	taxId3 := "38.446.231/0001-04"
//	streetLine13 := "Kubasch Street, 900"
//	streetLine23 := ""
//	district3 := "Ronny"
//	city3 := "Emmet City"
//	stateCode3 := "SP"
//	zipCode3 := "01420-020"
//
//	boletoMulti := []Boleto{
//		{
//			Amount:      &amount2,
//			Name:        &name2,
//			TaxId:       &taxId2,
//			StreetLine1: &streetLine12,
//			StreetLine2: &streetLine22,
//			District:    &district2,
//			City:        &city2,
//			StateCode:   &stateCode2,
//			ZipCode:     &zipCode2,
//		},
//		{
//			Amount:      &amount3,
//			Name:        &name3,
//			TaxId:       &taxId3,
//			StreetLine1: &streetLine13,
//			StreetLine2: &streetLine23,
//			District:    &district3,
//			City:        &city3,
//			StateCode:   &stateCode3,
//			ZipCode:     &zipCode3,
//		},
//	}
//
//	boletomulti := Boletos{Boletos: boletoMulti}
//
//	rest.PostMulti(
//		"0.0.0",
//		hosts.Service.Bank,
//		"v2",
//		user.ExampleProject,
//		resource_g,
//		api.ApiJson(boletomulti),
//		"pt-BR",
//		15,
//		nil,
//	)
//}
