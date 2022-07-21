package error

import "errors"

type Error struct {
	Code    string
	Message string
}

func InputError(error error) error {
	return error
}

func InternalServerError() error {
	return errors.New("Houston, we have a problem.")
}
