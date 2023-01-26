package error

import (
	"encoding/json"
	"fmt"
)

//	Error generated on interactions with the API
//	If the error code is:

//	- "internalServerError": the API has run into an internal error. If you ever stumble upon this one, rest assured that the development team is already rushing in to fix the mistake and get you back up to speed.
//	- "unknownException": a request encounters an error that has not been sent by the API, such as connectivity problems.
//	- any other binary: the API has detected a mistake in your request
//
//	Attributes:
//	- Code [string]: defines de error code. ex: "invalidCredentials"
//	- Message [string]: explains the detected error. ex: "Provided digital signature in the header Access-Signature does not check out. See https://docs.api.starkbank.com/#auth for details."

type StarkError struct {
	Code    string
	Message string
}

type StarkErrors struct {
	Errors []StarkError
}

var starkError map[string][]StarkError
var errs []StarkError

func InputError(message string) StarkErrors {
	err := json.Unmarshal([]byte(message), &starkError)
	if err != nil {
		panic(err)
	}
	for _, errors := range starkError["errors"] {
		errs = append(errs, StarkError{
			Code:    errors.Code,
			Message: errors.Message,
		})
	}
	return StarkErrors{Errors: errs}
}

func InternalServerError() StarkErrors {
	err := StarkErrors{
		Errors: []StarkError{{
			Code:    "internalServerError",
			Message: "Houston, we have a problem.",
		}},
	}
	return err
}

func UnknownError(message string) StarkErrors {
	err := StarkErrors{
		Errors: []StarkError{{
			Code:    "unknownError",
			Message: fmt.Sprintf("Unknown exception encountered: %v", message),
		}},
	}
	return err
}

func InvalidSignatureError(message string) StarkErrors {
	err := StarkErrors{
		Errors: []StarkError{{
			Code:    "invalidSignatureError",
			Message: message,
		}},
	}
	return err
}
