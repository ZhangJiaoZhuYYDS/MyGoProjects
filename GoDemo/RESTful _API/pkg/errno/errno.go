package errno

import "fmt"

type Errno struct {
	Code    int
	Message string
}

func (err Errno) Errno() string {
	return err.Message
}

type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}
func (errno *Errno) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s", errno.Code, errno.Message)
}

// New 新建定制错误
func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

// DecodeErr 解析定制错误
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
