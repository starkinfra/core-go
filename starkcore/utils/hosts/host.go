package hosts

type StarkHost struct {
	Infra string
	Bank  string
}

var Service = StarkHost{
	Bank:  "starkbank",
	Infra: "starkinfra",
}
