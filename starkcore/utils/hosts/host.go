package hosts

import "core-go/starkcore/utils/enum"

type StarkHoster interface {
	enum.Enumer
}

type StarkHost struct {
	Infra string
	Bank  string
}

var Service = StarkHost{
	Bank:  "starkbank",
	Infra: "starkinfra",
}
