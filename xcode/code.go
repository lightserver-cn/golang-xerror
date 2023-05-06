package xcode

import "net/http"

type Code struct {
	httpStatus int
	code       int
	message    string
}

func (c *Code) Code() int {
	return c.code
}

func (c *Code) HttpStatus() int {
	return c.httpStatus
}

func (c *Code) Message() string {
	return c.message
}

func NewCode(code int) Xcode {
	for _, c := range allCode {
		if c.code == code {
			return c
		}
	}

	if MinReservedCode <= code && code <= MaxReservedCode {
		return ErrCodeFailed
	}

	return &Code{
		httpStatus: http.StatusInternalServerError,
		code:       code,
		message:    ErrServCommon.Message(),
	}
}

func NewCodeMsg(code int, msg string) Xcode {

	if MinReservedCode <= code && code <= MaxReservedCode {
		return ErrCodeFailed
	}

	return &Code{
		httpStatus: http.StatusInternalServerError,
		code:       code,
		message:    msg,
	}
}
