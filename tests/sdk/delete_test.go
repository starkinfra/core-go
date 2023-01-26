package sdk

import (
	"fmt"
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	"testing"
)

func TestBoletoCancel(t *testing.T) {

	boleto, err := Boleto.Cancel("4537841761648640")
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Println(boleto)
}
