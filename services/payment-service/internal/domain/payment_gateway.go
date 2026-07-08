package domain

type MidtransGateway struct{
	PaymentUrl string
	GatewayOrderId string
}

type MidtransGatewayRepository interface{
	CreateTransaction(orderId string, amount float64) (string, string, error)
	GetPaymentStatus(gatewayOrderId string) (string, error)
}