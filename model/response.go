package model

import (
	"github.com/labstack/echo"
)

// Echo struct for response
type Echo struct {
	C echo.Context
}

// Response model json data
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// Response json string
func (ec *Echo) Response(code int, msg string, data interface{}) {
	ec.C.JSON(code, Response{
		Code:    code,
		Message: msg,
		Result:  data,
	})
}
