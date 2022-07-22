package resource

import (
	"core-go/starkcore/utils/subresource"
)

type Resource struct {
	Id          string
	Subresource subresource.Subresource
}

func (R Resource) name(resurce Resource) Resource {
	return resurce
}
