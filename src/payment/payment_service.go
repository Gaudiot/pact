package payment

import (
	"errors"
	"time"

	"github.com/pact/src/gateway"
	payment_gateway "github.com/pact/src/gateway/payment"
)

var (
	ErrInvalidAmount = errors.New("amount must be greater than zero")
)

type PaymentService struct {
	gateway      payment_gateway.PaymentGateway
	emailGateway gateway.EmailGateway
}

func NewPaymentService(paymentGateway payment_gateway.PaymentGateway, emailGateway gateway.EmailGateway) *PaymentService {
	return &PaymentService{
		paymentGateway,
		emailGateway,
	}
}

func (ps *PaymentService) CreatePix(amountInCents int) error {
	if amountInCents <= 0 {
		return ErrInvalidAmount
	}

	pixResponse, err := ps.gateway.CreatePixPayment(&payment_gateway.PixPaymentRequest{
		AmountInCents: amountInCents,
	})
	if err != nil {
		return err
	}

	err = ps.emailGateway.SendPixPaymentEmail(&gateway.PixSendPixPaymentEmailArgs{
		UserName:       "gaudiot1",
		AmountInCents:  pixResponse.AmountInCents,
		CompanyName:    "Fofolinda",
		ExpirationDate: time.Now().Add(time.Hour * 24 * 7),
		PixCode:        pixResponse.PaymentCode,
		PixQRCode:      pixResponse.PaymentQRCodeBase64,
		ReceiverEmail:  "gaudiot.dev@gmail.com",
	})
	if err != nil {
		return err
	}

	return nil
}

func (ps *PaymentService) CheckPixStatus(id string) error {
	_, err := ps.gateway.CheckPixPaymentStatus(&payment_gateway.CheckPixPaymentStatusRequest{Id: id})

	return err
}
