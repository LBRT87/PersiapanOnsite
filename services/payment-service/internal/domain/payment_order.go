package domain

type Order struct{
	OrderId string
	UserId string
	Amount string
	CreatedAt string
}

type OrderItem struct{
	OrderItemId string
	OrderId string
	CourseId string
	LecturerId string
	UserId string
}

type OrderRepository interface{
	Create(order Order) error
	FindByUserId(userId string) ([]Order, error)
	FindById(orderId string) (Order, error)
}

type OrderItemRepository interface{
	Create(orderItem OrderItem) error
	FindByLecturerId(lecturerId string) ([]OrderItem, error)
	CreateBatch(orderItems []OrderItem) error
}