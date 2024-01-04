package error

import (
	"fmt"
)

const (
	OKCode ErrCode = iota + 0
	// 注册相关
	ErrUserInfoNotProvided
	ErrParamsTypeErrorInServer
	ErrMySqlError
	ErrEmailAlreadyUsed
	ErrEmailNotProvided
	ErrPasswordNotProvided
	ErrGenderInvalid
	ErrUserTypeInvalid
	ErrPhoneNumberAlreadyUsed
	ErrGenerateTokenFailed
	ErrSetRedisFailed
	// 登录相关
	ErrorUserNotFound
	ErrorPasswordNotRight
)

type ErrCode uint32

// ErrorImpl
type ErrorImpl struct {
	ErrorCode ErrCode
	ErrorMsg  string
}

// New
func New(code ErrCode) ErrorImpl {
	return ErrorImpl{
	}
}

// Error
func (e ErrorImpl) Error() string {
	strFormat := `
    Error in api_data_query_server
    errorCode: %d
    errorMsg: %s
	`
	return fmt.Sprintf(strFormat, e.ErrorCode, e.ErrorMsg)
}
