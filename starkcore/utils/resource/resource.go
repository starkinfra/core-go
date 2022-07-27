package resource

import (
	"core-go/starkcore/utils/subresource"
	"fmt"
)

type Resource struct {
	Id          string
	Subresource subresource.Subresource
}

type Resourcezao interface {
	//GetNameResource() string
	//GetClass() string
}

func (R Resource) Resource() string {
	return fmt.Sprintf("%v[%v]", R.Subresource.Name, R.Id)
}
