package error

import (
	"errors"
	"fmt"
	"io"
)

type Error struct {
	Code    string
	Message string
}

func InputError(error io.ReadCloser) io.ReadCloser {
	return error
}

func InternalServerError() error {
	return errors.New("Houston, we have a problem. ")
}

func UnknownError() error {
	return errors.New("Unknown exception encountered: ")
}

func InvalidSignatureError(message string) string {
	return fmt.Sprintf("%v", message)
}
