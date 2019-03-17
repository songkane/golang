// Package request response
// Created by chenguolin 2018-11-17
package request

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.local.com/golang/go-common/logger"
	"gitlab.local.com/golang/go-common/trace"
	"gitlab.local.com/golang/go-common/uuid"
	golog "gitlab.local.com/golang/go-log"
	"gitlab.local.com/golang/go-prometheus"
)

var (
	errorCounterMonitor *prometheus.Counter
	succCounterMonitor  *prometheus.Counter
)

func init() {
	errorCounterMonitor = prometheus.NewCounter("internal_api_error_count_monitor",
		"internal api error count monitor", []string{"uri"})
	succCounterMonitor = prometheus.NewCounter("internal_api_succ_count_monitor",
		"internal api success count monitor", []string{"uri"})
}

const (
	// MetaKey meta
	MetaKey = "meta"
	// RequestIDKey request_id
	RequestIDKey = "request_id"
)

// EmptyStruct 空结构体定义
type EmptyStruct struct {
}

// Meta response meta定义
type Meta struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
	RequestID  string `json:"request_id"`
	RequestURI string `json:"request_uri"`
}

// Response 定义
type Response struct {
	Meta Meta        `json:"meta"`
	Body interface{} `json:"response"`
}

// APIError api响应失败
func APIError(c *gin.Context, errLog string) {
	reqID := GetRequestID(c)

	// build meta
	meta := Meta{
		Code:       500, //返回非0表示请求出错，需要业务配置暂定为500
		Message:    "",
		Error:      errLog,
		RequestID:  reqID,
		RequestURI: c.Request.RequestURI,
	}

	// output response json
	// default HTTP code 200
	outputResponseJSON(c, 200, meta, EmptyStruct{})

	// prometheus monitor
	labels := make(prometheus.Labels)
	labels["uri"] = c.Request.RequestURI
	errorCounterMonitor.Add(labels, 1)

	// print error log
	logger.Info(trace.WithTraceID(reqID), errLog,
		golog.Object("RequestParams", c.Request.Form),
		golog.Object("EndTime", time.Now().Unix()))

	// abort request default use 200 code
	c.AbortWithStatus(http.StatusOK)
}

// APISuccess api响应成功
func APISuccess(c *gin.Context, response interface{}) {
	reqID := GetRequestID(c)

	// build meta
	meta := Meta{
		Code:       0, //返回0表示请求正常
		Message:    "",
		Error:      "",
		RequestID:  reqID,
		RequestURI: c.Request.RequestURI,
	}

	// output response json
	// default HTTP code 200
	outputResponseJSON(c, 200, meta, response)

	// prometheus monitor
	labels := make(prometheus.Labels)
	labels["uri"] = c.Request.RequestURI
	succCounterMonitor.Add(labels, 1)

	// print error log
	logger.Info(trace.WithTraceID(reqID), "Successful process request ~",
		golog.Object("RequestParams", c.Request.Form),
		golog.Object("EndTime", time.Now().Unix()))
}

func outputResponseJSON(c *gin.Context, httpCode int, meta Meta, body interface{}) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(Response{
		Meta: meta,
		Body: body,
	})

	if err != nil {
		panic(err)
	}
	c.Set(MetaKey, meta)
	c.Data(httpCode, "application/json; charset=utf-8", buf.Bytes())
}

// GetRequestID 获取请求ID
func GetRequestID(c *gin.Context) string {
	// 先从context里面获取 如果有直接返回
	reqIDValue, isExist := c.Get(RequestIDKey)
	if isExist {
		return reqIDValue.(string)
	}

	// 生成RequestID
	reqID := uuid.NewUlid()
	c.Set(RequestIDKey, reqID)

	return reqID
}
