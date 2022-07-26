package checks

import (
	"core-go/starkcore/environment"
	"core-go/starkcore/utils/hosts"
	"fmt"
	"github.com/starkbank/ecdsa-go/ellipticcurve/privatekey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CheckEnvironment(env string) string {
	t := testing.T{}
	assert.Containsf(&t, environment.Environment, env, fmt.Sprintf("Select a valid environment %v", environment.Environment))
	return env
}

func CheckPrivateKey(pem string) string {
	if privatekey.FromPem(pem).Curve.Name == "secp256k1" {
		return pem
	} else {
		panic(fmt.Sprintf("Private-key must be valid secp256k1 ECDSA string in pem format"))
	}
}

func CheckLanguage(language string) string {
	t := testing.T{}
	acceptedLanguages := []string{"en-US", "pt-BR"}
	assert.Containsf(&t, language, acceptedLanguages, fmt.Sprintf("Language must be one from %v", acceptedLanguages))
	return language
}

func CheckHost(host string) string {
	if host == hosts.Service.Bank {
		service := "starkbank"
		return service
	}
	if host == hosts.Service.Infra {
		service := "starkinfra"
		return service
	}
	return ""
}
