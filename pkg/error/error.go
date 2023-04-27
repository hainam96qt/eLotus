package error

type XError struct {
	ErrorMessage string `json:"error_message"`
}

func (e *XError) Error() string {
	return e.ErrorMessage
}

func NewXError(s string) XError {
	return XError{
		ErrorMessage: s,
	}
}
