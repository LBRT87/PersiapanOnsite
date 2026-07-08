package usecase

import (
	"github.com/LBRT87/PersiapanOnsite/services/payment-service/config"
	"github.com/LBRT87/PersiapanOnsite/services/payment-service/internal/domain"
)

type PaymentUsecase struct {
	OrderRepo domain.OrderRepository
	OrderItemRepo domain.OrderItemRepository
	LecturerBalanceRepo domain.LecturerBalanceRepository
	WithdrawRepo domain.WithdrawalRepository
	Cfg *config.Config
}

