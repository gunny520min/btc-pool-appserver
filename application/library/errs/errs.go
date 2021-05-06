package errs

import (
	"errors"
	"fmt"

	"btc-pool-appserver/application/library/lang"
)

// 错误码
const (
	ErrnoSucc             = 0
	ErrnoTokenInvalid     = 20001
	ErrnoTokenError       = 20011
	ErrnoSystem           = 9000
	ErrnoSign             = 9005
	ErrnoInvalidParams    = 10001
	ErrnoInvalidAppId     = 10002
	ErrnoInvalidTimestamp = 10003
	ErrnoInvalidAppName   = 10004
	ErrnoLoginFail        = 1001
	ErrnoNeedRefreshToken = 1010
	ErrnoAuthNonce        = 10009
	ErrnoAmountTooSmall   = 955555
	ErrnoResultConvert    = 100001
)

// 一些系统常用错误
var (
	ApiErrLogin  *ErrApi
	ApiErrParams *ErrApi
	ApiErrSign   *ErrApi
	ApiErrSystem *ErrApi
)

type ErrApi struct {
	ErrNo  int
	ErrMsg string
	Data   interface{}
}

func (e *ErrApi) Error() string {
	return fmt.Sprintf("api err errno:%v errmsg:%s", e.ErrNo, e.ErrMsg)
}

// 初始化一些常用错误变量
func Init() {
	ApiErrLogin = NewApiError(ErrnoLoginFail, "", nil, "")
	ApiErrParams = NewApiError(ErrnoInvalidParams, "", nil, "")
	ApiErrSign = NewApiError(ErrnoSign, "", nil, "")
	ApiErrSystem = NewApiError(ErrnoSystem, "", nil, "")
}

// 判断一个错误是否是ApiErr类型
func UnwrapApiError(err error) *ErrApi {
	var apiErr *ErrApi
	if errors.As(err, &apiErr) {
		if apiErr != nil {
			return apiErr
		}
	}
	return nil
}

// 从lang配置读取错误码文案
func GetApiErrMessage(errNo int, language, suffix string) string {
	transKey := fmt.Sprintf("api.err_msg%s.%v", suffix, errNo)
	if language == "" {
		language = "en_US"
	}
	return lang.Trans(transKey, language)
}

// 构造一个api错误返回
func NewApiError(errNo int, errMsg string, data interface{}, language string) *ErrApi {
	if errMsg == "" {
		errMsg = GetApiErrMessage(errNo, language, "")
	}
	return &ErrApi{errNo, errMsg, data}
}

func NewApiErrorThird(errNo int, errMsg string, data interface{}, language string, suffix string) *ErrApi {
	if suffix == "" {
		suffix = "third"
	}
	tmpErrMsg := GetApiErrMessage(errNo, language, "_"+suffix)
	if tmpErrMsg != "" {
		errMsg = tmpErrMsg
	}
	return &ErrApi{errNo, errMsg, data}
}
