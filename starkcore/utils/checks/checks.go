package checks

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"core-go/starkcore/environment"
	"testing"
)

func CheckEnvironment(env string, t *testing.T) string{
	assert.Containsf(t,environment.Environment, env, "Select a valid environment" + for)
	return env
}

func CheckPrivateKey (pem string) string{

}

func CheckUser (user string) string{

}

func CheckLanguage (language string) string{

}

func CheckDateTimeOrDate (data string){

}

func CheckDateTime (user string){

}

func CheckDate (user string){

}

func CheckTimeDelta (user string){

}

func CheckDateTimeString (data string){

}