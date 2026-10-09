package api

import (
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"testing"
)

func TestLastNamePlural(t *testing.T) {
	cases := map[string]string{
		"AiSpeech":     "speeches",
		"Dish":         "dishes",
		"Tax":          "taxes",
		"Log":          "logs",
		"TransferLog":  "logs",
		"Transfer-Log": "logs",
		"Entity":       "entities",
		"Key":          "keys",
		"Address":      "address",
	}
	for name, expected := range cases {
		actual := api.LastNamePlural(map[string]string{"name": name})
		if actual != expected {
			t.Errorf("%s: expected %s, got %s", name, expected, actual)
		}
	}
}
