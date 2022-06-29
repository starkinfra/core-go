package environment

import "core-go/starkcore/utils/enum"

type Environmenter interface {
	enum.Enumer
}

type Environments struct {
	Sandbox    string
	Production string
}

var Environment = Environments{
	"sandbox",
	"production",
}
