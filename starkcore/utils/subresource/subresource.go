package subresource

import "fmt"

type Subresource struct {
	Name string
	Cls  interface{}
}

func (R Subresource) Subresource() string {
	return fmt.Sprintf("", R.Name)
}

func (R Subresource) ToString(resurce Resource) Resource {
	return resurce
}
