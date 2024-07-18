package error

type MustBindError struct {
	err error
}

func NewMustBindError(err error) *MustBindError {
	return &MustBindError{err: err}
}

func (m *MustBindError) Error() string {
	return m.err.Error()
}

func (m *MustBindError) Code() int {
	return 10001
}

func (m *MustBindError) Message() string {
	return m.err.Error()
}
