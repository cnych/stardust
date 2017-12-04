package errorsx

import (
	"bytes"
)

type MultiError []error

func (m MultiError) Error() string {
	if len(m) == 0 {
		return ""
	}
	buf := bytes.NewBufferString("")
	for _, err := range m {
		if err != nil {
			buf.WriteString(err.Error())
		} else {
			buf.WriteString("<nil>")
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func Multi(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	return MultiError(errs)
}

func SelectOrCompose(err1, err2 error) error {
	if err1 != nil && err2 != nil {
		return Multi([]error{err1, err2})
	}
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}
