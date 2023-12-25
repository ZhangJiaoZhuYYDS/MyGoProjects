package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred 结构体与请求体绑定."}

	ErrValidation = &Errno{Code: 20001, Message: "与结构体标签校验 failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "数据库 error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while 加密用户密码."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "用户不存在"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "密码不正确"}
)
