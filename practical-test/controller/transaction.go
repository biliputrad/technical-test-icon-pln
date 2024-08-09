package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	responseMessage "technical-test-icon-pln/practical-test/common/response-message"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	transactionDto "technical-test-icon-pln/practical-test/dto/transaction-dto"
	"technical-test-icon-pln/practical-test/service"
)

type transactionController struct {
	transactionService service.TransactionService
	pagination         paginate.Pagination
}

func NewTransactionController(transactionService service.TransactionService, pagination paginate.Pagination) *transactionController {
	return &transactionController{transactionService, pagination}
}

func (h *transactionController) Create(c *gin.Context) {
	var input transactionDto.CreateTransactionDto
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := responseMessage.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := responseMessage.GetResponse(http.StatusBadRequest, false, errorMessage, false)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := h.transactionService.Create(input)
	c.JSON(result.StatusCode, result)
}

func (h *transactionController) FindById(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		responseMessage.GetResponse(http.StatusBadRequest, false, "invalid id", false)
	}

	result := h.transactionService.FindById(id)
	c.JSON(result.StatusCode, result)
}

func (h *transactionController) FindAll(c *gin.Context) {
	pagination, search, _ := h.pagination.GetPagination(c)
	result := h.transactionService.FindAll(pagination, search)

	c.JSON(result.StatusCode, result)
}

func (h *transactionController) FindAllWithoutPagination(c *gin.Context) {
	result := h.transactionService.FindAllWithoutPagination()
	c.JSON(result.StatusCode, result)
}

func (h *transactionController) Delete(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		responseMessage.GetResponse(http.StatusBadRequest, false, "invalid id", false)
	}

	result := h.transactionService.Delete(id)
	c.JSON(result.StatusCode, result)
}
