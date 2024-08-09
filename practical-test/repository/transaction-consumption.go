package repository

import (
	"gorm.io/gorm"
	"technical-test-icon-pln/practical-test/common/constants"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	"technical-test-icon-pln/practical-test/model"
)

type TransactionConsumptionRepository interface {
	Create(transactionConsumption model.TransactionConsumption) (model.TransactionConsumption, error)
	Update(transactionConsumption model.TransactionConsumption) (model.TransactionConsumption, error)
	Delete(id int64) error
	FindAllWithQuery(pagination paginate.Pagination, query string) ([]model.TransactionConsumption, paginate.Pagination, error)
	FindById(id int64) (model.TransactionConsumption, error)
	FindAllWithoutPagination() ([]model.TransactionConsumption, error)
	FindAllWithoutQuery(pagination paginate.Pagination) ([]model.TransactionConsumption, paginate.Pagination, error)
}

type transactionConsumptionRepo struct {
	db *gorm.DB
}

func NewTransactionConsumptionRepository(db *gorm.DB) *transactionConsumptionRepo {
	return &transactionConsumptionRepo{db}
}

func (r *transactionConsumptionRepo) Create(transactionConsumption model.TransactionConsumption) (model.TransactionConsumption, error) {
	err := r.db.Create(&transactionConsumption).Error

	return transactionConsumption, err
}

func (r *transactionConsumptionRepo) Update(transactionConsumption model.TransactionConsumption) (model.TransactionConsumption, error) {
	err := r.db.Save(&transactionConsumption).Error

	return transactionConsumption, err
}

func (r *transactionConsumptionRepo) Delete(id int64) error {
	err := r.db.Where(constants.ById, id).Delete(&model.TransactionConsumption{}).Error

	return err
}

func (r *transactionConsumptionRepo) FindAllWithQuery(pagination paginate.Pagination, query string) ([]model.TransactionConsumption, paginate.Pagination, error) {
	var countries []model.TransactionConsumption

	err := r.db.Scopes(paginate.Paginate(&countries, &pagination, r.db)).Where(query).Find(&countries).Error

	return countries, pagination, err
}

func (r *transactionConsumptionRepo) FindById(id int64) (model.TransactionConsumption, error) {
	var result model.TransactionConsumption

	err := r.db.Where(constants.ById, id).First(&result).Error

	return result, err
}

func (r *transactionConsumptionRepo) FindAllWithoutPagination() ([]model.TransactionConsumption, error) {
	var countries []model.TransactionConsumption

	err := r.db.Find(&countries).Error

	return countries, err
}

func (r *transactionConsumptionRepo) FindAllWithoutQuery(pagination paginate.Pagination) ([]model.TransactionConsumption, paginate.Pagination, error) {
	var countries []model.TransactionConsumption

	err := r.db.Scopes(paginate.Paginate(&countries, &pagination, r.db)).Find(&countries).Error

	return countries, pagination, err
}
