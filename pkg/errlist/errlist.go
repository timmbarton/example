package errlist

import "github.com/timmbarton/errors"

var (
	ErrInternal = errs.New(errs.ErrCodeInternal, 10_0001, "some internal error")
)
