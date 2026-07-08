package domain

type CustomerInfo struct{
	Name string
	Email string
}

type MidtransGatewayRepository interface{
	CreateTransaction(orderId string, amount float64, customer CustomerInfo) (string, string, error)
	GetPaymentStatus(gatewayOrderId string) (string, error)
}