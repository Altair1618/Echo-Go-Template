package errors

type AppError struct {
	Code    int
	Message string
	Err     error
	Fields  map[string]string
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}
