package response

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Error
	logLevel int
	err      error
	aborted  bool
}

func newResponse() *Response {
	return &Response{}
}

func (e *Response) defaultData(data interface{}) interface{} {
	if data == nil {
		data = Error{e.Code, e.Msg}
	}
	return data
}

func (e *Response) coverData(data Error) interface{} {
	if e.Code != 0 {
		data.Code = e.Code
	}

	if e.Msg != "" {
		data.Msg = e.Msg
	}
	return data
}

func (e *Response) finishContext(c *gin.Context, statusCode int, data interface{}) {
	if e.aborted {
		c.AbortWithStatusJSON(statusCode, data)
		return
	}
	c.JSON(statusCode, data)
}

func (e *Response) WithCode(c int) *Response {
	e.Code = c
	return e
}

func (e *Response) WithMsg(s ...interface{}) *Response {
	e.Msg = fmt.Sprint(s)
	return e
}

func (e *Response) WithCodeAndMsg(ec Error) *Response {
	e.Code = ec.Code
	e.Msg = ec.Msg
	return e
}

func (e *Response) Log(err error) *Response {
	caller, file, line, ok := runtime.Caller(1)
	fmt.Print(caller, file, line, ok, err)
	return e
}

func (e *Response) Abort() *Response {
	e.aborted = true
	return e
}

func (e *Response) InternalErr(c *gin.Context) {
	e.aborted = true
	e.finishContext(c, http.StatusInternalServerError, e.coverData(InternalError))
}

func (e *Response) ForbiddenErr(c *gin.Context) {
	e.aborted = true
	e.finishContext(c, http.StatusForbidden, e.coverData(Forbidden))
}

func (e *Response) NotFoundErr(c *gin.Context) {
	e.aborted = true
	e.finishContext(c, http.StatusForbidden, e.coverData(NotFound))
}

func (e *Response) UnauthorizedErr(c *gin.Context) {
	e.aborted = true
	e.finishContext(c, http.StatusUnauthorized, e.coverData(Unauthorized))
}

func (e *Response) BadRequestErr(c *gin.Context) {
	e.aborted = true
	e.finishContext(c, http.StatusBadRequest, e.coverData(BadRequest))
}

func (e *Response) Success(c *gin.Context, data interface{}) {
	data = e.defaultData(data)

	e.finishContext(c, http.StatusOK, data)
}

//WithCode 会将改errCode绑定在返回的response中，不可用于Success返回
func WithCode(c int) *Response { r := newResponse(); r.Code = c; return r }

//WithMsg 会将这些信息包装到返回的response.msg中，不可用于Success返回
func WithMsg(s ...interface{}) *Response { r := newResponse(); r.Msg = fmt.Sprint(s); return r }

func WithCodeAndMsg(e Error) *Response {
	r := newResponse()
	r.Code = e.Code
	r.Msg = e.Msg
	return r
}

func WithMsgLog(s ...interface{}) *Response {
	r := newResponse()
	r.Msg = fmt.Sprint(s)
	r.Log(errors.New(fmt.Sprint(s)))
	return r
}

//Log 会将error信息打印出来，包括调用的位置、时间等信息
func Log(err error) *Response {
	r := newResponse()
	fmt.Print(runtime.Caller(1))
	fmt.Println("err:", err)
	return r
}

func Success(ctx *gin.Context, data interface{}) { newResponse().Success(ctx, data) }

func InternalErr(c *gin.Context) { newResponse().InternalErr(c) }

func ForbiddenErr(c *gin.Context) { newResponse().ForbiddenErr(c) }

func NotFoundErr(c *gin.Context) { newResponse().NotFoundErr(c) }

func UnauthorizedErr(c *gin.Context) { newResponse().UnauthorizedErr(c) }

func BadRequestErr(c *gin.Context) { newResponse().BadRequestErr(c) }
