package error

import (
	"net/http"
	"server/pkg/codes"

	"github.com/gin-gonic/gin"
)

func SendErrJSON(msg string, args ...interface{}) {
	if len(args) == 0 {
		panic("no *gin.Context")
	}
	var c *gin.Context
	var errCode = codes.ERROR

	if len(args) == 1 {
		theCtx, ok := args[0].(*gin.Context)
		if !ok {
			panic("No More *gin.Context")
		}
		c = theCtx
	} else if len(args) == 2 {
		theErrNo, ok := args[0].(int)
		if !ok {
			panic("errCode not Define")
		}
		errCode = theErrNo
		theCtx, ok := args[1].(*gin.Context)
		if !ok {
			panic("no *gin.Context")
		}
		c = theCtx
	}

	c.JSON(http.StatusOK, gin.H{
		"status": errCode,
		"msg":    msg,
	})
	c.Abort()
}
