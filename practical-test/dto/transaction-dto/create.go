package transaction_dto

import "time"

type CreateTransactionDto struct {
	BookingDate time.Time
	OfficeName  string
	StartTime   time.Time
	EndTime     time.Time
	Participant int64
	RoomName    string
}
