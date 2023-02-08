package handler

import (
	"net/http"
	"strconv"

	"github.com/fauziagustian/delos-aqua/internal/dto"
	"github.com/fauziagustian/delos-aqua/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetFarm(c *gin.Context) {

	query := &dto.RequestQuery{}
	err := c.ShouldBindQuery(query)
	if err != nil {
		if err != nil {
			errors := utils.FormatValidationError(err)
			response := utils.ErrorResponse("get farm failed", http.StatusUnprocessableEntity, errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	query = dto.FormatQuery(query)

	farm, err := h.farmService.GetFarm(query)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("get farm failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	totalFarm, err := h.farmService.CountFarm()
	if err != nil {
		response := utils.ErrorResponse("get farm failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedFarm := dto.FormatFarms(farm)
	metadata := utils.Metadata{Resource: "farms", TotalAll: int(totalFarm), TotalNow: len(farm), Page: query.Page, Limit: query.Limit, Sort: query.Sort}
	response := utils.ResponseWithPagination("get farm success", http.StatusOK, formattedFarm, metadata)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateFarm(c *gin.Context) {

	input := &dto.FarmRequestBody{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("Create Farm failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	farm, err := h.farmService.CreateFarm(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("Create Farm failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedFarm := dto.FormatFarm(farm)
	response := utils.SuccessResponse("Create Farm Successfully", http.StatusOK, formattedFarm)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateFarm(c *gin.Context) {
	input := &dto.FarmRequestBody{}
	getId := c.Param("id")

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("Updated Farm failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	id, err := strconv.Atoi(getId)
	if err != nil {
		return
	}

	farm, err := h.farmService.UpdateFarm(input, id)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("Updated Farm failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedFarm := dto.FormatFarm(farm)
	response := utils.SuccessResponse("Updated Farm Successfully", http.StatusOK, formattedFarm)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteFarm(c *gin.Context) {
	getId := c.Param("id")
	id, err := strconv.Atoi(getId)
	if err != nil {
		return
	}

	_, err = h.farmService.DeleteFarm(id)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("Deleted Farm failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	response := utils.SuccessResponse("Deleted Farm id : "+getId+" Successfully", http.StatusOK, "")
	c.JSON(http.StatusOK, response)
}
