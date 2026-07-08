package usecase

import (
	"context"
	"errors"

	coursepb "github.com/LBRT87/PersiapanOnsite/contracts/course-service/gen"
	"github.com/LBRT87/PersiapanOnsite/services/payment-service/config"
	"github.com/LBRT87/PersiapanOnsite/services/payment-service/internal/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentUsecase struct {
	OrderRepo domain.OrderRepository
	OrderItemRepo domain.OrderItemRepository
	LecturerBalanceRepo domain.LecturerBalanceRepository
	WithdrawRepo domain.WithdrawalRepository
	Cfg *config.Config
}

func NewPaymentUseCase(orderRepo domain.OrderRepository, orderItemRepo domain.OrderItemRepository, lecturerBalanceRepo domain.LecturerBalanceRepository,
withdrawalRepo domain.WithdrawalRepository, cfg *config.Config) *PaymentUsecase{
	return &PaymentUsecase{OrderRepo: orderRepo, OrderItemRepo: orderItemRepo, LecturerBalanceRepo: lecturerBalanceRepo, WithdrawRepo: withdrawalRepo, Cfg: cfg}
}

func (u *PaymentUsecase) dialInsecure(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func (u *PaymentUsecase) Checkout(userId string, lecturerId string, coursesId []string) (string, error){
	if len(coursesId) == 0{
		return "", errors.New("Theres should be atleast 1 course")
	}
	orderItems := make([]domain.OrderItem, 0, len(coursesId))
	// for _, courseId := range coursesId{
		
	// }
	return "", nil
}

func (u *PaymentUsecase) IsEnrolled(userId, courseId string) (bool, error){
	conn, err := u.dialInsecure(u.Cfg.CourseServiceAddr);
	if err != nil{
		return false, err
	}
	defer conn.Close()
	resp, err := coursepb.NewCourseServiceClient(conn).GetEnrollment(context.Background(), &coursepb.GetEnrollmentRequest{UserId : userId, CourseId : courseId})
	if err != nil{
		return false, err
	}
	return resp.Enrolled, nil
}