package memory

import "errors"

var (
	ErrIllegalValue = errors.New("Illegal value (nil or cache.TheMissing)? ")
)
