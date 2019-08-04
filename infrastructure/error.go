package infrastructure

import (
	"fmt"
	"golang.org/x/xerrors"
)

const (
	Unknown = iota
	NotFound
	CovertDocumentRefToStruct
)

type NotFoundError struct {
	msg  string
	code int
}

func (e *NotFoundError) Error() string {
	return e.msg
}

type CovertDocumentRefToStructError struct {
	msg  string
	code int
}

func (e *CovertDocumentRefToStructError) Error() string {
	return e.msg
}

func Errorf(err error, errorCode int) error {
	switch errorCode {
	case NotFound:
		return &NotFoundError{
			msg:  fmt.Sprintf("%+v", xerrors.Errorf("get data failed from firestore document ref: %w", err)),
			code: NotFound,
		}
	case CovertDocumentRefToStruct:
		return &CovertDocumentRefToStructError{
			msg:  fmt.Sprintf("%+v", xerrors.Errorf("resource not found: %w", err)),
			code: CovertDocumentRefToStruct,
		}
	default:
		return xerrors.Errorf("unknown error: %w", err)
	}
}
