package errors

type InvalidToken struct {
	Err error
}

func (e *InvalidToken) Error() string {
	return "invalid token"
}

type NoContextFound struct {
	Err error
}

func (e *NoContextFound) Error() string {
	return "no context found"
}

type InvalidRequestType struct {
	Err error
}

func (e *InvalidRequestType) Error() string {
	return "invalid request type"
}

type InvalidPhoneNumber struct {
	Err error
}

func (e *InvalidPhoneNumber) Error() string {
	return "invalid phone number"
}
