package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserAuthenticate(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello maotuan")
}
