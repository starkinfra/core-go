package checks

import (
	"fmt"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/privatekey"
	"github.com/starkinfra/core-go/starkcore/environment"
	"reflect"
	"strings"
)

func CheckEnvironment(env string) string {
	var acceptedEnvironments []string
	v := reflect.ValueOf(environment.Environments)
	for i := 0; i < v.NumField(); i++ {
		acceptedEnvironments = append(acceptedEnvironments, v.Field(i).String())
		if env == v.Field(i).Interface() {
			return env
		}
	}
	panic(fmt.Sprintf("Select a valid environment: %v", strings.Join(acceptedEnvironments, " or ")))
}

func CheckPrivateKey(pem string) string {
	if privatekey.FromPem(pem).Curve.Name == "secp256k1" {
		return pem
	}
	panic("Private-key must be valid secp256k1 ECDSA string in pem format")
}

func CheckLanguage(language string) string {
	acceptedLanguages := []string{"en-US", "pt-BR"}
	for _, validLanguages := range acceptedLanguages {
		if validLanguages == language {
			return language
		}
	}
	panic(fmt.Sprintf("Language must be one from %v", strings.Join(acceptedLanguages, ", ")))
}
