package handler

import (
	"net/http"

	"github.com/fauziagustian/delos-aqua/internal/dto"
	"github.com/fauziagustian/delos-aqua/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserAgent(c *gin.Context) {
	userAgent, err := h.userAgentService.GetUserAgent()
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("get user agent failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedUserAgents := dto.FormatUserAgentsCount(userAgent)
	response := utils.SuccessResponse("get user agent success", http.StatusOK, formattedUserAgents)
	c.JSON(http.StatusOK, response)
}
