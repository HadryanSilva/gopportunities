package handler

import (
	"fmt"
	"github.com/HadryanSilva/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"error": message,
		"code":  code,
	})
}

func sendSuccess(ctx *gin.Context, code int, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("Operation from handler: %s successfully", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
