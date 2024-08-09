package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	responseMessage "technical-test-icon-pln/practical-test/common/response-message"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	transactionConsumptionDto "technical-test-icon-pln/practical-test/dto/transaction-consumption-dto"
	"technical-test-icon-pln/practical-test/service"
)

type transactionConsumptionController struct {
	transactionConsumptionService service.TransactionConsumptionService
	pagination                    paginate.Pagination
}

func NewTransactionConsumptionController(transactionConsumptionService service.TransactionConsumptionService, pagination paginate.Pagination) *transactionConsumptionController {
	return &transactionConsumptionController{transactionConsumptionService, pagination}
}

func (h *transactionConsumptionController) Create(c *gin.Context) {
	var input transactionConsumptionDto.CreateTransactionConsumptionDto
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := responseMessage.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := responseMessage.GetResponse(http.StatusBadRequest, false, errorMessage, false)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := h.transactionConsumptionService.Create(input)
	c.JSON(result.StatusCode, result)
}

func (h *transactionConsumptionController) FindById(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		responseMessage.GetResponse(http.StatusBadRequest, false, "invalid id", false)
	}

	result := h.transactionConsumptionService.FindById(id)
	c.JSON(result.StatusCode, result)
}

func (h *transactionConsumptionController) FindAll(c *gin.Context) {
	pagination, search, _ := h.pagination.GetPagination(c)
	result := h.transactionConsumptionService.FindAll(pagination, search)

	c.JSON(result.StatusCode, result)
}

func (h *transactionConsumptionController) FindAllWithoutPagination(c *gin.Context) {
	result := h.transactionConsumptionService.FindAllWithoutPagination()
	c.JSON(result.StatusCode, result)
}

func (h *transactionConsumptionController) Delete(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		responseMessage.GetResponse(http.StatusBadRequest, false, "invalid id", false)
	}

	result := h.transactionConsumptionService.Delete(id)
	c.JSON(result.StatusCode, result)
}
