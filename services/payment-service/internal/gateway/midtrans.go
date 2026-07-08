package gateway

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/LBRT87/PersiapanOnsite/services/payment-service/internal/domain"
)

type MidtransGateway struct {
	ServerKey    string
	ClientKey    string
	IsProduction bool
}

func NewMidtransGateway(serverKey, ClientKey string, isProduction bool) *MidtransGateway {
	return &MidtransGateway{ServerKey: serverKey, ClientKey: ClientKey, IsProduction: isProduction}
}

type MidtransSnapRequest struct {
	MidtransUserInfo  MidtransUserInfo  `json:"midtrans_user_info"`
	MidtransOrderInfo MidtransOrderInfo `json:"midtrans_order_info"`
}

type MidtransUserInfo struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
}

type MidtransOrderInfo struct {
	OrderId     string  `json:"order_id"`
	GrossAmount float64 `json:"gross_amount"`
}

type MidtransSnapResponse struct {
	Token       string `json:"token"`
	RedirectUrl string `json:"redirect_url"`
}

const (
	TaskQueue    = "checkout-task-queue"
	WorkflowName = "CheckoutCourseWorfklow"
)

func (g *MidtransGateway) SnapURL() string {
	if g.IsProduction {
		return "https://app.midtrans.com/snap/v1/transactions"
	}
	return "https://app.sandbox.midtrans.com/snap/v1/transactions"
}

func (g *MidtransGateway) StatusURL(orderID string) string {
	if g.IsProduction {
		return "https://api.midtrans.com/v2/" + orderID + "/status"
	}
	return "https://api.sandbox.midtrans.com/v2/" + orderID + "/status"
}

func (g *MidtransGateway) CreateTransaction(orderId string, GrossAmount float64, customer domain.CustomerInfo) (string, string, error){
	reqBody := MidtransSnapRequest{
		MidtransUserInfo{FirstName: customer.Name, Email: customer.Email},
		MidtransOrderInfo{OrderId: orderId, GrossAmount: GrossAmount},
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", "",  err
	}
	req, err := http.NewRequest(http.MethodPost, g.SnapURL(), bytes.NewReader(body))
	if err != nil{
		return "", "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(g.ServerKey+":")))
	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		return "", "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return "", "", err
	}
	var snapResp MidtransSnapResponse
	if err := json.NewDecoder(resp.Body).Decode(&snapResp); err != nil{
		return "", "", err
	}
	return snapResp.RedirectUrl, orderId, nil
}