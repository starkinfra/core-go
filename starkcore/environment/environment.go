package environment

type Environments struct {
	Sandbox    string
	Production string
}

var Environment = Environments{
	"sandbox",
	"production",
}
