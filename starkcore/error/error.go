package error

type Error struct {
	Code    string
	Message string
}

type InputError struct {
	errors Error
}

type InternalServerError struct {
	Code    string
	Message string
}

type UnkownError struct {
	Code    string
	Message string
}

type InvalidSignatureError struct {
	Code    string
	Message string
}
