package model

import "time"

type Consumption struct {
	ID                      int64
	CreatedAt               time.Time
	Name                    string
	MaxPrice                int64
	TransactionConsumptions []TransactionConsumption `gorm:"foreignKey:ConsumptionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
