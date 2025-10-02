package log

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/core-go/tests/utils"
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

func Get(id string) (Log, Error.StarkErrors) {
	get, err := rest.GetId(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoletoLog,
		id,
		nil,
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

func Query(params map[string]interface{}) (chan Log, chan Error.StarkErrors) {
	logChannel := make(chan Log)
	streamChannel, errChannel := rest.GetStream(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoletoLog,
		params,
	)
	go func() {
		for were := range streamChannel {
			wereByte, _ := json.Marshal(were)
			err := json.Unmarshal(wereByte, &log)
			if err != nil {
				print(err)
			}
			logChannel <- log
		}
		close(logChannel)
	}()
	return logChannel, errChannel
}

func Page(params map[string]interface{}) ([]Log, string, Error.StarkErrors) {
	page, cursor, err := rest.GetPage(
		utils.SdkVersion,
		hosts.Bank,
		utils.ApiVersion,
		utils.Language,
		utils.Timeout,
		User.ExampleProjectBank,
		utils.ResourceBoletoLog,
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
