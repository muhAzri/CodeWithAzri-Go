package response

import (
	"github.com/gin-gonic/gin"
)

func BuildData(payload interface{}) *Response {
	return &Response{Data: payload}
}

func Respond(code int, payload interface{}, ctx *gin.Context) {
	response := &Response{
		Data: payload,
	}
	ctx.JSON(code, response)
}

func RespondError(code int, err error, ctx *gin.Context) {
	response := &Response{
		Error: map[string]string{"error": err.Error()},
		Meta: Meta{
			Message: "Error",
			Code:    code,
			Status:  "error",
		},
	}
	ctx.JSON(code, response)
}

func RespondErrorMessage(code int, msg string, ctx *gin.Context) {
	response := &Response{
		Error: map[string]string{"error": msg},
		Meta: Meta{
			Message: "Error",
			Code:    code,
			Status:  "error",
		},
	}
	ctx.JSON(code, response)
}
