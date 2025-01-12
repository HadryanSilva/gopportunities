package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}
