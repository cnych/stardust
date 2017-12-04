package errorsx

import (
	"errors"
	"fmt"
)

func AsError(v interface{}) error {
	if v == nil {
		return nil
	}
	switch err := v.(type) {
	case error:
		return err
	case string:
		return errors.New(err)
	default:
		return fmt.Errorf("%v", err)
	}
}
