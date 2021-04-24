package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/serializer"
)

// errors
var (
	ErrLoginRequired = errors.New("请先登录")
	ErrBadRequest    = errors.New("错误的请求")
	ErrEmptyParam    = errors.New("参数不可为空")
	ErrInvalidParam  = errors.New("参数不合法")
	
	ErrEmptyValue = errors.New("empty value")
)

// Action structure
type Action struct {
	Method        Method
	Action        ActionFunc
	LoginRequired bool
	
	handler gin.HandlerFunc
}

// GetHandler for request handler
func (a *Action) GetHandler() gin.HandlerFunc {
	return a.handler
}

// NewAction creates new action
func NewAction(method Method, handler ActionFunc, loginRequired bool) *Action {
	a := Action{Method: method, Action: handler, LoginRequired: loginRequired}
	a.handler = func(c *gin.Context) {
		// init context
		ctx := Context{c}
		
		if a.LoginRequired {
			abortError(c, ErrLoginRequired)
			return
		}
		
		r, err := a.Action(&ctx)
		if err != nil {
			abortError(c, err)
			return
			
		}
		c.JSON(http.StatusOK, serializer.Response{Code: serializer.CodeSuccess, Data: r})
	}
	return &a
}

func abortError(c *gin.Context, err error) {
	c.Status(http.StatusInternalServerError)
	abortWithError(c, http.StatusInternalServerError, serializer.CodeParamErr, err.Error())
	
}

func abortWithError(c *gin.Context, status, code int, info interface{}) {
	c.AbortWithStatusJSON(status, serializer.Response{Code: code, Data: info})
}
