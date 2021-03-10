package log

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

// init logger
func InitLogger(logWriter io.Writer) {
	textFormatter := new(logrus.JSONFormatter)
	textFormatter.TimestampFormat = "2006-01-02 15:04:05.000"
	//textFormatter.DisableColors = true
	logrus.SetFormatter(textFormatter)
	logrus.SetOutput(logWriter)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func ContextInfo(c *gin.Context, args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"trace_id": GetRequestID(c),
	}).Info(args...)
}

func ContextWithFields(c *gin.Context, info string, data map[string]interface{}) {
	var fields = logrus.Fields{
		"trace_id": GetRequestID(c),
	}

	if data != nil {
		for k, v := range data {
			fields[k] = v
		}
	}

	logrus.WithFields(fields).Info(info)
}

func ContextError(c *gin.Context, args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"uri":      c.Request.URL.Path,
		"trace_id": GetRequestID(c),
	}).Error(args...)
}

func GetRequestID(c *gin.Context) string {
	v, ok := c.Get("_reqID")
	if !ok {
		return ""
	}
	if reqID, ok := v.(string); ok {
		return reqID
	}
	return ""
}
