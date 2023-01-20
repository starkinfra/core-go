package log

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils/boleto"
	User "github.com/starkinfra/core-go/tests/utils/user"
	"time"
)

type Log struct {
	Id      string        `json:",omitempty"`
	Boleto  boleto.Boleto `json:",omitempty"`
	Errors  []string      `json:",omitempty"`
	Type    string        `json:",omitempty"`
	Created *time.Time    `json:",omitempty"`
}

var log Log
var logs []Log
var resourceBoletoLog = map[string]string{"name": "BoletoLog"}

type BoletoHolmes struct {
	BoletoId string     `json:",omitempty"`
	Tags     string     `json:",omitempty"`
	Id       string     `json:",omitempty"`
	Status   string     `json:",omitempty"`
	Result   string     `json:",omitempty"`
	Created  *time.Time `json:",omitempty"`
	Updated  *time.Time `json:",omitempty"`
}

type HolmesLog struct {
	Id      string       `json:",omitempty"`
	Holmes  BoletoHolmes `json:",omitempty"`
	Errors  []string     `json:",omitempty"`
	Type    string       `json:",omitempty"`
	Created *time.Time   `json:",omitempty"`
}

var ResourceHolmesLog = map[string]string{"name": "BoletoHolmesLog"}

func Get(id string) (Log, Error.StarkErrors) {
	get, err := rest.GetId(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		resourceBoletoLog,
		id,
	)
	if err.Errors != nil {
		return Log{}, err
	}
	unmarshalError := json.Unmarshal(get, &log)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return log, err
}

func Query(params map[string]interface{}) ([]Log, Error.StarkErrors) {
	query, err := rest.GetStream(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		resourceBoletoLog,
		params,
	)
	if err.Errors != nil {
		return []Log{}, err
	}
	unmarshalError := json.Unmarshal(query, &logs)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return logs, err
}

func Page(params map[string]interface{}) ([]Log, string, Error.StarkErrors) {
	page, cursor, err := rest.GetPage(
		User.SdkVersion,
		hosts.Bank,
		User.ApiVersion,
		User.Language,
		User.Timeout,
		User.ExampleProjectBank,
		resourceBoletoLog,
		params,
	)
	if err.Errors != nil {
		return []Log{}, "", err
	}
	unmarshalError := json.Unmarshal(page, &logs)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}
	return logs, cursor, err
}
