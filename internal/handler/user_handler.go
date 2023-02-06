package handler

import (
	"net/http"

	"github.com/fauziagustian/delos-aqua/internal/dto"
	"github.com/fauziagustian/delos-aqua/internal/models"
	"github.com/fauziagustian/delos-aqua/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Profile(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	formattedUser := dto.FormatUserDetail(user)
	response := utils.SuccessResponse("show profile success", http.StatusOK, formattedUser)
	c.JSON(http.StatusOK, response)
}
