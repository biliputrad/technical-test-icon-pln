package model

import "time"

type Transaction struct {
	ID                      int64
	BookingDate             time.Time
	OfficeName              string
	StartTime               time.Time
	EndTime                 time.Time
	Participant             int64
	RoomName                string
	TransactionConsumptions []TransactionConsumption `gorm:"foreignKey:TransactionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
