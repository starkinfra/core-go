package error

import (
	"fmt"
)

type StarkError interface {
}

type Errors struct {
	Code    string
	Message string
	Content string
}

func Error(errors Errors) string {
	return fmt.Sprintf("%c : %m", errors.Code, errors.Message)
}

func InputErrors() string {
	return fmt.Sprintf()
}

func InternalServerError() string {
	return fmt.Sprintf("Houston, we have a problem")
}

func UnkownError() string {
	return fmt.Sprintf("Unkown exception encountered: %m", Message)
}

func InvalidSignatureError() string {
	return fmt.Sprintf("")
}
