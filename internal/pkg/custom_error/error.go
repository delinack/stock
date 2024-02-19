package custom_error

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrUnavailableStock = errors.New("stock is unavailable")
	ErrUnexpected       = errors.New("unexpected error")
	ErrNullValue        = errors.New("wrong params in request: null value")
	ErrExceededValue    = errors.New("wrong params in request: exceeded value")
	ErrWithTransaction  = errors.New("transaction error")
)
