package handler

import (
	"fmt"
	"github.com/HadryanSilva/gopportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOpeningHandler
// @Base path /api/v1
// @Summary Find an opening
// @Description Find a job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} GetOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id %s not found", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "get-opening", opening)
}
