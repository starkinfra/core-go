package checks

import (
	"core-go/starkcore/environment"
	"core-go/starkcore/user/organization"
	"core-go/starkcore/user/users"
	"core-go/starkcore/utils/hosts"
	"fmt"
	"github.com/starkbank/ecdsa-go/ellipticcurve/privatekey"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func CheckEnvironment(env string, t *testing.T) string {
	message := fmt.Sprintf("Select a valid environment %m", strings.Join(environment.Environment, ","))
	assert.Containsf(t, environment.Environment, env, message)
	return env
}

func CheckPrivateKey(pem string) string {
	if (privatekey.FromPem(pem).Curve.Name == "secp256k1"){

	} else {
		return fmt.Sprintf("Private-key must be valid secp256k1 ECDSA string in pem format")
	}
	return pem
}

func CheckUser(user users.Users) users.Users {
	vari, _ := fmt.Printf("%T", user)
	variable := string(vari)
	if (variable != "user.User"){
		fmt.Printf("A user is required to access our API. Check our README: https://github.com/starkinfra/core-python/")
	} 
	return user
}

func CheckAccess (users users.Users) string{
	
	if (users. ){
		organization.AccessId()
	} else if (us) {

	}
	return ""
}

func CheckLanguage(language string) string {
	acceptedLanguages := []string{"en-US", "pt-BR"}

	message := fmt.Sprintf("Language must be one from %a", acceptedLanguages)
	assert.Containsf(t, language, acceptedLanguages, message)
	return language
}

func CheckHost(host string) string{
	if (host == hosts.Service.Bank){
		service := "starkbank"
		return service
	}
	if (host == hosts.Service.Infra) {
		service := "starkinfra"
		return service
	}
	return ""
}

func CheckDateTimeOrDate(data string) {

}

func CheckDateTime(user string) {

}

func CheckDate(user string) {

}

func CheckTimeDelta(data string) {
	if &data == nil {
		return ""
	}

	if data.

}

func CheckDateTimeString(data string) {
	data = string(data)

}
