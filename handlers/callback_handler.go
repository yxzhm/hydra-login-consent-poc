package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CallbackHandler(ctx *gin.Context) {
	code := ctx.Request.URL.Query().Get("code")
	ctx.String(http.StatusOK, code)
}
