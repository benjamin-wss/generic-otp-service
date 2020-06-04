package dto

type ApiErrorGeneric struct {
	HttpStatus int
	Error      error
}

func (e *ApiErrorGeneric) Unwrap() error { return e.Error }
