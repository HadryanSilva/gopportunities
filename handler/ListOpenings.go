package handler

import (
	"github.com/HadryanSilva/gopportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListOpeningsHandler
// @Base path /api/v1
// @Summary Find an opening
// @Description Find a job opening
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	var openings []schemas.Opening

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to list openings")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-openings", openings)
}
