package handler

import (
	"net/http"

	"github.com/fauziagustian/delos-aqua/internal/dto"
	"github.com/fauziagustian/delos-aqua/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	input := &dto.RegisterRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("register failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.CreateUser(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("register failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	token, err := h.jwtService.GenerateToken(int(newUser.UserId))
	if err != nil {
		response := utils.ErrorResponse("register failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedLogin := dto.FormatLogin(newUser, token)
	response := utils.SuccessResponse("register success", http.StatusOK, formattedLogin)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Login(c *gin.Context) {
	input := &dto.LoginRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("login failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.authService.Attempt(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("login failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	token, err := h.jwtService.GenerateToken(int(loggedinUser.UserId))
	if err != nil {
		response := utils.ErrorResponse("login failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedLogin := dto.FormatLogin(loggedinUser, token)
	response := utils.SuccessResponse("login success", http.StatusOK, formattedLogin)
	c.JSON(http.StatusOK, response)
}
