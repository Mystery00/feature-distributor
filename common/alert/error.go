package alert

import (
	"errors"
	"fmt"
)

func Error(code Code) error {
	return &ErrorCode{Code: code}
}

func Convert(err error) *Code {
	var e *ErrorCode
	if errors.As(err, &e) {
		return &e.Code
	}
	return nil
}

type ErrorCode struct {
	Code Code
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("%d", e.Code)
}
