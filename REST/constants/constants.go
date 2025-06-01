package constants

import "errors"

var (
	ErrBadRequest  = errors.New("BadRequest")
	ErrInternal    = errors.New("Internal")
	ErrWrongFormat = errors.New("WrongFormat")
	ErrNotAllowed  = errors.New("NotAllowed")
)
