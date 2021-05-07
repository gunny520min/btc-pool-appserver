package output

import (
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/log"
	"bytes"
	encjson "encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpJsonResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"msg,omitempty"`
	TraceId   string      `json:"trace_id,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	TimeStamp int64       `json:"timestamp"`
}

// 支持任何err类型，如果是我们定义的api err，则按api err输出，否则一律输出500系统错误
// 建议都用该方法进行错误输出, 有日志记录
func ShowErr(c *gin.Context, e error) {
	var ginError *gin.Error

	if errors.As(e, &ginError) { // 如果已经是gin.Error类型，则取原始的error，并且不在写进c.Errors
		e = ginError.Err
	} else {
		_ = c.Error(e) // 写进gin.context 一些middleware例如sentry可以捕捉错误
	}

	if apiE := errs.UnwrapApiError(e); apiE != nil {
		errApi(c, apiE)
	} else {
		errApi(c, errs.ApiErrSystem)
	}

	log.ContextError(c, "end: ", e)
}

func json(c *gin.Context, httpCode, errNo int, data interface{}, errMsg string) {
	if data == nil {
		data = struct{}{}
	}
	c.JSON(httpCode, gin.H{
		"code":      errNo,
		"msg":       errMsg,
		"timestamp": time.Now().Unix(),
		"trace_id":  log.GetRequestID(c),
		"data":      data,
	})
}

func GenerateJsonResponse(code int, message, traceId string, data interface{}) ([]byte, error) {

	var httpResp *HttpJsonResponse
	if code == 0 {
		// sucess
		httpResp = &HttpJsonResponse{
			Code:      code,
			Message:   message,
			TimeStamp: time.Now().Unix(),
			TraceId:   traceId,
			Data:      data,
		}
	} else {
		httpResp = &HttpJsonResponse{
			Code:      code,
			Message:   message,
			TimeStamp: time.Now().Unix(),
			TraceId:   traceId,
		}
	}

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := encjson.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(httpResp)
	return bf.Bytes(), nil
}

func WriteHttpJsonResponse(context *gin.Context, code, errNo int, data interface{}, message string) {

	context.Status(http.StatusOK)
	rawData, err := GenerateJsonResponse(errNo, message, log.GetRequestID(context), data)
	if err != nil {
		panic(err)
	}

	header := context.Writer.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}

	context.Writer.Write(rawData)
	context.Writer.Flush()
}

func SuccGray(c *gin.Context, data interface{}) {
	errNo := errs.ErrnoSucc
	if code, exist := c.Get("_forceOutputCode"); exist {
		if codeForce, ok := code.(int); ok {
			errNo = codeForce
		}
	}
	//json(c, http.StatusOK, errNo, data, "success")
	WriteHttpJsonResponse(c, http.StatusOK, errNo, data, "success")
	log.ContextWithFields(c, "end_succ", map[string]interface{}{
		"data": data,
	})
}

func Succ(c *gin.Context, data interface{}) {
	errNo := errs.ErrnoSucc
	if code, exist := c.Get("_forceOutputCode"); exist {
		if codeForce, ok := code.(int); ok {
			errNo = codeForce
		}
	}
	json(c, http.StatusOK, errNo, data, "success")
	log.ContextWithFields(c, "end_succ", map[string]interface{}{
		"data": data,
	})
}

func SuccList(c *gin.Context, list interface{}) {
	data := gin.H{
		"list":        list,
	}
	Succ(c, data)
}

func errorComm(c *gin.Context, errNo int, data interface{}, errMsg string) {
	json(c, http.StatusOK, errNo, data, errMsg)
	c.AbortWithStatus(http.StatusOK)
}

func errApi(c *gin.Context, e *errs.ErrApi) {
	errorComm(c, e.ErrNo, e.Data, e.ErrMsg)
}

// 注意 httpcode=500
func errSystem(c *gin.Context) {
	json(c, http.StatusInternalServerError, errs.ErrnoSystem, nil, "An unknown error")
	c.AbortWithStatus(http.StatusInternalServerError)
}

// 直接输出一个json字符串
func SuccJsonStr(c *gin.Context, str string) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	res := fmt.Sprintf("{\"err_no\":%v,\"err_msg\":\"%s\",\"data\":%s}", errs.ErrnoSucc, "succ", str)

	c.String(http.StatusOK, res)
}

func SuccPage(c *gin.Context, list interface{}, count int) {
	limit := c.GetInt("limit")
	var pageTotal int
	if count%limit > 0 {
		pageTotal = (count / limit) + 1
	} else {
		pageTotal = count / limit
	}
	data := gin.H{
		"list":        list,
		"page":        c.GetInt("page"),
		"page_size":   limit,
		"page_total":  pageTotal,
		"total_count": count,
	}
	Succ(c, data)
}
