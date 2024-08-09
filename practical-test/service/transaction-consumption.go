package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"technical-test-icon-pln/practical-test/common/constants"
	responseMessage "technical-test-icon-pln/practical-test/common/response-message"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	transactionConsumptionDto "technical-test-icon-pln/practical-test/dto/transaction-consumption-dto"
	"technical-test-icon-pln/practical-test/model"
	"technical-test-icon-pln/practical-test/repository"
)

type TransactionConsumptionService interface {
	Create(dto transactionConsumptionDto.CreateTransactionConsumptionDto) responseMessage.Response
	FindById(id int64) responseMessage.Response
	FindAll(pagination paginate.Pagination, search string) responseMessage.ResponsePaginate
	FindAllWithoutPagination() responseMessage.Response
	Delete(id int64) responseMessage.Response
}

type transactionConsumptionService struct {
	transactionConsumptionRepository repository.TransactionConsumptionRepository
}

func NewTransactionConsumptionService(transactionConsumptionRepository repository.TransactionConsumptionRepository) *transactionConsumptionService {
	return &transactionConsumptionService{transactionConsumptionRepository}
}

func (s *transactionConsumptionService) Create(dto transactionConsumptionDto.CreateTransactionConsumptionDto) responseMessage.Response {
	data := dtoToModelsCreateTransactionConsumption(dto)
	result, err := s.transactionConsumptionRepository.Create(data)
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    constants.ResponseCreated,
		Data:       result,
	}
}

func dtoToModelsCreateTransactionConsumption(dto transactionConsumptionDto.CreateTransactionConsumptionDto) model.TransactionConsumption {
	return model.TransactionConsumption{
		TransactionId: dto.TransactionId,
		ConsumptionId: dto.ConsumptionId,
	}
}

func (s *transactionConsumptionService) FindById(id int64) responseMessage.Response {
	result, err := s.transactionConsumptionRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusNotFound,
				Success:    false,
				Message:    fmt.Errorf("transactionConsumption with id %d not found", id).Error(),
				Data:       nil,
			}

		} else {
			return responseMessage.Response{
				StatusCode: http.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
				Data:       nil,
			}
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
	}
}

func (s *transactionConsumptionService) FindAll(pagination paginate.Pagination, search string) responseMessage.ResponsePaginate {
	var result responseMessage.ResponsePaginate
	if search != "" {
		result = s.findAllWithFilter(pagination, search)
	} else {
		result = s.findAllWithoutQuery(pagination)
	}

	return result
}

func (s *transactionConsumptionService) findAllWithFilter(pagination paginate.Pagination, search string) responseMessage.ResponsePaginate {
	query := "name ILIKE '%" + search + "%' OR max_price::VARCHAR(255) ILIKE '%" + search + "%'"

	result, paginateResult, err := s.transactionConsumptionRepository.FindAllWithQuery(pagination, query)
	if err != nil {
		return responseMessage.ResponsePaginate{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
			Pagination: nil,
		}
	}

	return responseMessage.ResponsePaginate{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
		Pagination: paginateResult,
	}
}

func (s *transactionConsumptionService) findAllWithoutQuery(pagination paginate.Pagination) responseMessage.ResponsePaginate {
	result, paginateResult, err := s.transactionConsumptionRepository.FindAllWithoutQuery(pagination)
	if err != nil {
		return responseMessage.ResponsePaginate{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
			Pagination: nil,
		}
	}

	return responseMessage.ResponsePaginate{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
		Pagination: paginateResult,
	}
}

func (s *transactionConsumptionService) FindAllWithoutPagination() responseMessage.Response {
	result, err := s.transactionConsumptionRepository.FindAllWithoutPagination()
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
	}
}

func (s *transactionConsumptionService) Delete(id int64) responseMessage.Response {
	_, err := s.transactionConsumptionRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusNotFound,
				Success:    false,
				Message:    fmt.Errorf("transactionConsumption with id %d not found", id).Error(),
				Data:       nil,
			}

		} else {
			return responseMessage.Response{
				StatusCode: http.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
				Data:       nil,
			}
		}
	}

	err = s.transactionConsumptionRepository.Delete(id)
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       true,
	}
}
