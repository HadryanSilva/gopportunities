package handler

import (
	"github.com/HadryanSilva/gopportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningsHandler(ctx *gin.Context) {
	var openings []schemas.Opening

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to list openings")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-openings", openings)
}
