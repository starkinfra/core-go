package environment

type Environment struct {
	Sandbox    string
	Production string
}

var Environments = Environment{
	"sandbox",
	"production",
}
