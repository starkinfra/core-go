package subresource

type Subresource struct {
	Name string
}

func (R Resource) name(resurce Resource) Resource {
	return resurce
}
