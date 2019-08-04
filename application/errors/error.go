package errors

import "golang.org/x/xerrors"

const (
	Unknown = iota
	ResourceNotFound
)

type ResourceNotFoundError struct {
	msg string
}

func (e *ResourceNotFoundError) Error() string {
	return e.msg
}

func Errorf(err error, errorCode int) error {
	switch errorCode {
	case ResourceNotFound:
		return &ResourceNotFoundError{xerrors.Errorf("resource not found: %+v", err).Error()}
	default:
		return xerrors.Errorf("unknown error: %w", err)
	}
}

func UnWrap(err error, count int) error {
	unwrapped := err
	for {
		unwrapped = xerrors.Unwrap(unwrapped)
		count--
		if count <= 0 {
			break
		}
	}
	return unwrapped
}
