package url

import (
	"core-go/starkcore/utils/api"
	"fmt"
)

func UrlEncode(params interface{}) string {
	params = api.CastJsonToApiFormat(params)
	if params != nil {
		responseQuery := fmt.Sprintf("?", params)
		return (responseQuery)
	} else {
		return ""
	}
}
