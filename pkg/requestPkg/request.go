package requestPkg

import "github.com/gin-gonic/gin"

func GetUserID(ctx *gin.Context) string {
	userID := ctx.MustGet("UserId").(string)
	return userID
}

func GetURLParam(ctx *gin.Context, key string) string {
	return ctx.Param(key)
}

func GetQueryParam(ctx *gin.Context, key string) string {
	return ctx.Query(key) // shortcut for c.Request.URL.Query().Get("lastname")
}
