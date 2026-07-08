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

}

type OrderItemRepository interface{
	
}