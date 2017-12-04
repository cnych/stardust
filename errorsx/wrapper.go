package errorsx

import (
	"fmt"
)

type WrapError struct {
	Err     error
	Message string
}

func (w *WrapError) Error() string {
	if w == nil {
		return ""
	}
	return fmt.Sprintf("%s cause: %s", w.Message, w.Err.Error())
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &WrapError{err, msg}
}
