package handler

import (
	"fmt"
	"github.com/HadryanSilva/gopportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOpeningHandler
// @Base path /api/v1
// @Summary Delete an opening
// @Description Delete a new job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	// Find opening by id
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id %s not found", id))
		return
	}
	// Delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Failed to delete opening with id %s", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "delete-opening", opening)
}
