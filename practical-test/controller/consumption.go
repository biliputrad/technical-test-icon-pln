package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	responseMessage "technical-test-icon-pln/practical-test/common/response-message"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	consumptionDto "technical-test-icon-pln/practical-test/dto/consumption-dto"
	"technical-test-icon-pln/practical-test/service"
)

type consumptionController struct {
	consumptionService service.ConsumptionService
	pagination         paginate.Pagination
}

func NewConsumptionController(consumptionService service.ConsumptionService, pagination paginate.Pagination) *consumptionController {
	return &consumptionController{consumptionService, pagination}
}

func (h *consumptionController) Create(c *gin.Context) {
	var input consumptionDto.CreateConsumptionDto
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := responseMessage.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := responseMessage.GetResponse(http.StatusBadRequest, false, errorMessage, false)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := h.consumptionService.Create(input)
	c.JSON(result.StatusCode, result)
}

func (h *consumptionController) FindById(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		responseMessage.GetResponse(http.StatusBadRequest, false, "invalid id", false)
	}

	result := h.consumptionService.FindById(id)
	c.JSON(result.StatusCode, result)
}

func (h *consumptionController) FindAll(c *gin.Context) {
	pagination, search, _ := h.pagination.GetPagination(c)
	result := h.consumptionService.FindAll(pagination, search)

	c.JSON(result.StatusCode, result)
}

func (h *consumptionController) FindAllWithoutPagination(c *gin.Context) {
	result := h.consumptionService.FindAllWithoutPagination()
	c.JSON(result.StatusCode, result)
}

func (h *consumptionController) Delete(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		responseMessage.GetResponse(http.StatusBadRequest, false, "invalid id", false)
	}

	result := h.consumptionService.Delete(id)
	c.JSON(result.StatusCode, result)
}
