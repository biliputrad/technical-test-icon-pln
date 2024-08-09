package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"technical-test-icon-pln/practical-test/common/constants"
	responseMessage "technical-test-icon-pln/practical-test/common/response-message"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	consumptionDto "technical-test-icon-pln/practical-test/dto/consumption-dto"
	"technical-test-icon-pln/practical-test/model"
	"technical-test-icon-pln/practical-test/repository"
	"time"
)

type ConsumptionService interface {
	Create(dto consumptionDto.CreateConsumptionDto) responseMessage.Response
	FindById(id int64) responseMessage.Response
	FindAll(pagination paginate.Pagination, search string) responseMessage.ResponsePaginate
	FindAllWithoutPagination() responseMessage.Response
	Delete(id int64) responseMessage.Response
}

type consumptionService struct {
	consumptionRepository repository.ConsumptionRepository
}

func NewConsumptionService(consumptionRepository repository.ConsumptionRepository) *consumptionService {
	return &consumptionService{consumptionRepository}
}

func (s *consumptionService) Create(dto consumptionDto.CreateConsumptionDto) responseMessage.Response {
	data := dtoToModelsCreateConsumption(dto)
	result, err := s.consumptionRepository.Create(data)
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

func dtoToModelsCreateConsumption(dto consumptionDto.CreateConsumptionDto) model.Consumption {
	return model.Consumption{
		CreatedAt: time.Now(),
		Name:      dto.Name,
		MaxPrice:  dto.MaxPrice,
	}
}

func (s *consumptionService) FindById(id int64) responseMessage.Response {
	result, err := s.consumptionRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusNotFound,
				Success:    false,
				Message:    fmt.Errorf("consumption with id %d not found", id).Error(),
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

func (s *consumptionService) FindAll(pagination paginate.Pagination, search string) responseMessage.ResponsePaginate {
	var result responseMessage.ResponsePaginate
	if search != "" {
		result = s.findAllWithFilter(pagination, search)
	} else {
		result = s.findAllWithoutQuery(pagination)
	}

	return result
}

func (s *consumptionService) findAllWithFilter(pagination paginate.Pagination, search string) responseMessage.ResponsePaginate {
	query := "name ILIKE '%" + search + "%' OR max_price::VARCHAR(255) ILIKE '%" + search + "%'"

	result, paginateResult, err := s.consumptionRepository.FindAllWithQuery(pagination, query)
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

func (s *consumptionService) findAllWithoutQuery(pagination paginate.Pagination) responseMessage.ResponsePaginate {
	result, paginateResult, err := s.consumptionRepository.FindAllWithoutQuery(pagination)
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

func (s *consumptionService) FindAllWithoutPagination() responseMessage.Response {
	result, err := s.consumptionRepository.FindAllWithoutPagination()
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

func (s *consumptionService) Delete(id int64) responseMessage.Response {
	_, err := s.consumptionRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusNotFound,
				Success:    false,
				Message:    fmt.Errorf("consumption with id %d not found", id).Error(),
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

	err = s.consumptionRepository.Delete(id)
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
