package transaction_dto

import "time"

type Result struct {
	BookingDate     time.Time
	OfficeName      string
	StartTime       time.Time
	EndTime         time.Time
	ListConsumption []ListConsumptionResult
	Participants    int64
	RoomName        string
	Id              int64
}

type ListConsumptionResult struct {
	Name string
}
