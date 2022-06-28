package error

import (
	"fmt"
)

type error interface {
	Error() string
}

type Error struct {
	Code    string
	Message string
	Content string
}

func Error() string {
	return fmt.Sprintf("%c : %m", e.Code, e.Message)
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
