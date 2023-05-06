package xcode

type Xcode interface {
	Code() int
	HttpStatus() int
	Message() string
}
