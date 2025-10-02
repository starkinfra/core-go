package sdk

import (
	Boleto "github.com/starkinfra/core-go/tests/utils/boleto"
	"github.com/starkinfra/core-go/tests/utils/examples"
	"testing"
)

func TestBoletoCreateAndCancel(t *testing.T) {
	createdBoleto, err := Boleto.Create(examples.ExampleBoleto())
	if err.Errors != nil {
		t.Errorf("err: %s", err.Errors)
	}

	_, err = Boleto.Cancel(createdBoleto[0].Id)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
}
