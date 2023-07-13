package xerr

import "fmt"

// 常用通用固定错误

type CodeError struct {
	errCode uint32
	errMsg  string
}

// GetErrorCode 返回给前端的错误码
func (e *CodeError) GetErrorCode() uint32 {
	return e.errCode
}

// GetErrorMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrorMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrorCode: %d, ErrorMsg: %s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{
		errCode: errCode,
		errMsg:  MapErrMsg(errCode),
	}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{
		errCode: SERVER_COMMON_ERROR,
		errMsg:  errMsg,
	}
}

func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}
