package application

import (
	"fmt"
	"golang.org/x/xerrors"
)

const (
	Unknown = iota
	ResourceNotFound
	Internal
)

type ResourceNotFoundError struct {
	msg  string
	code int
}

func (e *ResourceNotFoundError) Error() string {
	return e.msg
}

type InternalError struct {
	msg  string
	code int
}

func (e *InternalError) Error() string {
	return e.msg
}

func Errorf(err error, errorCode int) error {
	switch errorCode {
	case ResourceNotFound:
		return &ResourceNotFoundError{
			msg:  fmt.Sprintf("%+v", xerrors.Errorf("resource not found: %w", err)),
			code: ResourceNotFound,
		}
	case Internal:
		return &InternalError{
			msg:  fmt.Sprintf("%+v", xerrors.Errorf("internal error: %w", err)),
			code: Internal,
		}
	default:
		return xerrors.Errorf("unknown error: %w", err)
	}
}

func UnWrap(err error) error {
	return xerrors.Unwrap(err)
}
