package errorx

const defaultCode = 1001

// CodeError 错误信息
type CodeError struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

// CodeErrorResponse 代码错误的返回封装信息
type CodeErrorResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

//func NewDefaultError(msg string) error {
//	return NewCodeError(defaultCode, msg)
//}

func NewDefaultError(code int) error {
	return NewCodeError(code, MapErrMsg(code))
}

func (err *CodeError) Error() string {
	return err.Msg
}
func (err *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    err.Code,
		Msg:     err.Msg,
		Success: err.Success,
	}
}
