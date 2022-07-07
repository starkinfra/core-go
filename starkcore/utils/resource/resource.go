package resource

import "core-go/starkcore/utils/subresource"

type Resource struct {
	Id          string
	Subresource subresource.Subresource
}
