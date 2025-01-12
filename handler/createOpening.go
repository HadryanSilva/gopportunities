package handler

import (
	"github.com/HadryanSilva/gopportunities/config"
	"github.com/HadryanSilva/gopportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOpeningHandler
// @Base path /api/v1
// @Summary Create an opening
// @Description Create a new job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request Body"
// @Success 201 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("Validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("Failed to create opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Failed to create opening")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-opening", opening)
}

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetPostgres()
}
