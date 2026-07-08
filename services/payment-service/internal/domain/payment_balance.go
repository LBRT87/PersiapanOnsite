package domain

import "time"

type LecturerBalance struct {
	BalanceId string `gorm:"primaryKey;default:gen_random_uuid()"`
	LecturerId string
	TotalWithdrawn float64
	TotalEarnings float64
	UpdateAt time.Time
}

type Withdrawal struct{
	WithdrawalId string `gorm:"primaryKey;default:gen_random_uuid()"`
	LecturerId   string
	Amount       float64
	WithdrawAt   time.Time
}

type LecturerBalanceRepository interface{
	Create(lecturerBalance LecturerBalance) error
	CreditEarnings(lecturerId string, amount float64) error
	DebitFromWithdrawal(lecturerId string, amount float64) error
	GetBalance(lecturerId string) (LecturerBalance, error)
}

type WithdrawalRepository interface{
	Create(withdraw Withdrawal) error
	FindByLecturerId(lecturerId string) ([]Withdrawal, error)
}