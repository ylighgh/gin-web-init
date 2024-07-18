package error

type CodeError struct {
	err error
}

func NewCodeError(err error) *CodeError {
	return &CodeError{err: err}
}

func (m *CodeError) Error() string {
	return m.err.Error()
}

func (m *CodeError) Code() int {
	return 400
}

func (m *CodeError) Message() string {
	return m.err.Error()
}
