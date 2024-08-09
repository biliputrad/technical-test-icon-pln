package repository

import (
	"gorm.io/gorm"
	"technical-test-icon-pln/practical-test/common/constants"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	"technical-test-icon-pln/practical-test/model"
)

type ConsumptionRepository interface {
	Create(consumption model.Consumption) (model.Consumption, error)
	Update(consumption model.Consumption) (model.Consumption, error)
	Delete(id int64) error
	FindAllWithQuery(pagination paginate.Pagination, query string) ([]model.Consumption, paginate.Pagination, error)
	FindById(id int64) (model.Consumption, error)
	FindAllWithoutPagination() ([]model.Consumption, error)
	FindAllWithoutQuery(pagination paginate.Pagination) ([]model.Consumption, paginate.Pagination, error)
}

type consumptionRepo struct {
	db *gorm.DB
}

func NewConsumptionRepository(db *gorm.DB) *consumptionRepo {
	return &consumptionRepo{db}
}

func (r *consumptionRepo) Create(consumption model.Consumption) (model.Consumption, error) {
	err := r.db.Create(&consumption).Error

	return consumption, err
}

func (r *consumptionRepo) Update(consumption model.Consumption) (model.Consumption, error) {
	err := r.db.Save(&consumption).Error

	return consumption, err
}

func (r *consumptionRepo) Delete(id int64) error {
	err := r.db.Where(constants.ById, id).Delete(&model.Consumption{}).Error

	return err
}

func (r *consumptionRepo) FindAllWithQuery(pagination paginate.Pagination, query string) ([]model.Consumption, paginate.Pagination, error) {
	var countries []model.Consumption

	err := r.db.Scopes(paginate.Paginate(&countries, &pagination, r.db)).Where(query).Find(&countries).Error

	return countries, pagination, err
}

func (r *consumptionRepo) FindById(id int64) (model.Consumption, error) {
	var result model.Consumption

	err := r.db.Where(constants.ById, id).First(&result).Error

	return result, err
}

func (r *consumptionRepo) FindAllWithoutPagination() ([]model.Consumption, error) {
	var countries []model.Consumption

	err := r.db.Find(&countries).Error

	return countries, err
}

func (r *consumptionRepo) FindAllWithoutQuery(pagination paginate.Pagination) ([]model.Consumption, paginate.Pagination, error) {
	var countries []model.Consumption

	err := r.db.Scopes(paginate.Paginate(&countries, &pagination, r.db)).Find(&countries).Error

	return countries, pagination, err
}
