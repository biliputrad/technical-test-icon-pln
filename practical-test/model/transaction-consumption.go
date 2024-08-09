package model

type TransactionConsumption struct {
	ID            int64
	TransactionId int64
	Transaction   Transaction `gorm:"foreignKey:TransactionId;references:ID"`
	ConsumptionId int64
	Consumption   Consumption `gorm:"foreignKey:ConsumptionId;references:ID"`
}
