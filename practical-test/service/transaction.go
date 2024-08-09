package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"technical-test-icon-pln/practical-test/common/constants"
	responseMessage "technical-test-icon-pln/practical-test/common/response-message"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	transactionDto "technical-test-icon-pln/practical-test/dto/transaction-dto"
	"technical-test-icon-pln/practical-test/model"
	"technical-test-icon-pln/practical-test/repository"
)

type TransactionService interface {
	Create(dto transactionDto.CreateTransactionDto) responseMessage.Response
	FindById(id int64) responseMessage.Response
	FindAll(pagination paginate.Pagination, search string) responseMessage.ResponsePaginate
	FindAllWithoutPagination() responseMessage.Response
	Delete(id int64) responseMessage.Response
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) *transactionService {
	return &transactionService{transactionRepository}
}

func (s *transactionService) Create(dto transactionDto.CreateTransactionDto) responseMessage.Response {
	data := dtoToModelsCreateTransaction(dto)
	result, err := s.transactionRepository.Create(data)
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

func dtoToModelsCreateTransaction(dto transactionDto.CreateTransactionDto) model.Transaction {
	return model.Transaction{
		BookingDate: dto.BookingDate,
		OfficeName:  dto.OfficeName,
		StartTime:   dto.StartTime,
		EndTime:     dto.EndTime,
		Participant: dto.Participant,
		RoomName:    dto.RoomName,
	}
}

func (s *transactionService) FindById(id int64) responseMessage.Response {
	result, err := s.transactionRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusNotFound,
				Success:    false,
				Message:    fmt.Errorf("transaction with id %d not found", id).Error(),
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

func (s *transactionService) FindAll(pagination paginate.Pagination, search string) responseMessage.ResponsePaginate {
	var result responseMessage.ResponsePaginate
	if search != "" {
		result = s.findAllWithFilter(pagination, search)
	} else {
		result = s.findAllWithoutQuery(pagination)
	}

	return result
}

func (s *transactionService) findAllWithFilter(pagination paginate.Pagination, search string) responseMessage.ResponsePaginate {
	query := "office_name ILIKE '%" + search + "%' OR participant::VARCHAR(255) ILIKE '%" + search + "%' OR room_name ILIKE '%" + search + "%'"

	transactionData, paginateResult, err := s.transactionRepository.FindAllWithQuery(pagination, query)
	if err != nil {
		return responseMessage.ResponsePaginate{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
			Pagination: nil,
		}
	}

	var result []transactionDto.Result
	for _, data := range transactionData {
		var resultConsumption []transactionDto.ListConsumptionResult
		for _, consumptionData := range data.TransactionConsumptions {
			resultConsumption = append(resultConsumption, transactionDto.ListConsumptionResult{
				Name: consumptionData.Consumption.Name,
			})
		}

		result = append(result, transactionDto.Result{
			Id:              data.ID,
			BookingDate:     data.BookingDate,
			OfficeName:      data.OfficeName,
			StartTime:       data.StartTime,
			EndTime:         data.EndTime,
			Participants:    data.Participant,
			RoomName:        data.RoomName,
			ListConsumption: resultConsumption,
		})
	}

	return responseMessage.ResponsePaginate{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
		Pagination: paginateResult,
	}
}

func (s *transactionService) findAllWithoutQuery(pagination paginate.Pagination) responseMessage.ResponsePaginate {
	result, paginateResult, err := s.transactionRepository.FindAllWithoutQuery(pagination)
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

func (s *transactionService) FindAllWithoutPagination() responseMessage.Response {
	result, err := s.transactionRepository.FindAllWithoutPagination()
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

func (s *transactionService) Delete(id int64) responseMessage.Response {
	_, err := s.transactionRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusNotFound,
				Success:    false,
				Message:    fmt.Errorf("transaction with id %d not found", id).Error(),
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

	err = s.transactionRepository.Delete(id)
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
