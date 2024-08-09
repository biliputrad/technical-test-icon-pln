package repository

import (
	"gorm.io/gorm"
	"technical-test-icon-pln/practical-test/common/constants"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	"technical-test-icon-pln/practical-test/model"
)

type TransactionRepository interface {
	Create(transaction model.Transaction) (model.Transaction, error)
	Update(transaction model.Transaction) (model.Transaction, error)
	Delete(id int64) error
	FindAllWithQuery(pagination paginate.Pagination, query string) ([]model.Transaction, paginate.Pagination, error)
	FindById(id int64) (model.Transaction, error)
	FindAllWithoutPagination() ([]model.Transaction, error)
	FindAllWithoutQuery(pagination paginate.Pagination) ([]model.Transaction, paginate.Pagination, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepo {
	return &transactionRepo{db}
}

func (r *transactionRepo) Create(transaction model.Transaction) (model.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *transactionRepo) Update(transaction model.Transaction) (model.Transaction, error) {
	err := r.db.Save(&transaction).Error

	return transaction, err
}

func (r *transactionRepo) Delete(id int64) error {
	err := r.db.Where(constants.ById, id).Delete(&model.Transaction{}).Error

	return err
}

func (r *transactionRepo) FindAllWithQuery(pagination paginate.Pagination, query string) ([]model.Transaction, paginate.Pagination, error) {
	var countries []model.Transaction

	err := r.db.Scopes(paginate.Paginate(&countries, &pagination, r.db)).Where(query).Find(&countries).Error

	return countries, pagination, err
}

func (r *transactionRepo) FindById(id int64) (model.Transaction, error) {
	var result model.Transaction

	err := r.db.Where(constants.ById, id).First(&result).Error

	return result, err
}

func (r *transactionRepo) FindAllWithoutPagination() ([]model.Transaction, error) {
	var countries []model.Transaction

	err := r.db.Find(&countries).Error

	return countries, err
}

func (r *transactionRepo) FindAllWithoutQuery(pagination paginate.Pagination) ([]model.Transaction, paginate.Pagination, error) {
	var countries []model.Transaction

	err := r.db.Scopes(paginate.Paginate(&countries, &pagination, r.db)).Find(&countries).Error

	return countries, pagination, err
}
