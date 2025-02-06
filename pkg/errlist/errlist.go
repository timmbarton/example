package errlist

import "github.com/timmbarton/errors"

var (
	ErrStopTimeout = errs.New(errs.ErrCodeInternal, 10_0001, "stop timeout")
)
