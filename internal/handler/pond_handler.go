package handler

import (
	"net/http"
	"strconv"

	"github.com/fauziagustian/delos-aqua/internal/dto"
	"github.com/fauziagustian/delos-aqua/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPond(c *gin.Context) {
	getUserId, _ := c.Get("userid")
	userId := getUserId.(int)

	query := &dto.RequestQuery{}
	err := c.ShouldBindQuery(query)
	if err != nil {
		if err != nil {
			errors := utils.FormatValidationError(err)
			response := utils.ErrorResponse("get pond failed", http.StatusUnprocessableEntity, errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	query = dto.FormatQuery(query)

	pond, err := h.pondService.GetPond(query, userId)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("get pond failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	totalPond, err := h.pondService.CountPond()
	if err != nil {
		response := utils.ErrorResponse("get pond failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedPond := dto.FormatPonds(pond)
	metadata := utils.Metadata{Resource: "ponds", TotalAll: int(totalPond), TotalNow: len(pond), Page: query.Page, Limit: query.Limit, Sort: query.Sort}
	response := utils.ResponseWithPagination("get pond success", http.StatusOK, formattedPond, metadata)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) CreatePond(c *gin.Context) {
	getUserId, _ := c.Get("userid")
	userId := getUserId.(int)

	input := &dto.PondRequestBody{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("Create Pond failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	pond, err := h.pondService.CreatePond(input, userId)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("Create Pond failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedPond := dto.FormatPond(pond)
	response := utils.SuccessResponse("Create Pond Successfully", http.StatusOK, formattedPond)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdatePond(c *gin.Context) {
	getUserId, _ := c.Get("userid")
	userId := getUserId.(int)

	input := &dto.PondRequestBody{}
	getId := c.Param("id")

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("Updated Pond failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	id, err := strconv.Atoi(getId)
	if err != nil {
		return
	}

	pond, err := h.pondService.UpdatePond(input, id, userId)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("Updated Pond failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedPond := dto.FormatPond(pond)
	response := utils.SuccessResponse("Updated Pond Successfully", http.StatusOK, formattedPond)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) DeletePond(c *gin.Context) {
	getUserId, _ := c.Get("userid")
	userId := getUserId.(int)

	getId := c.Param("id")
	id, err := strconv.Atoi(getId)
	if err != nil {
		return
	}

	_, err = h.pondService.DeletePond(id, userId)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("Deleted Pond failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	response := utils.SuccessResponse("Deleted Pond id : "+getId+" Successfully", http.StatusOK, "")
	c.JSON(http.StatusOK, response)
}
