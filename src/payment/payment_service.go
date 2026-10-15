package payment

import (
	"errors"
	"time"

	"github.com/pact/src/gateway"
)

var (
	ErrInvalidAmount = errors.New("amount must be greater than zero")
)

type PaymentService struct {
	gateway      PaymentGateway
	emailGateway gateway.EmailGateway
}

func NewPaymentService(paymentGateway PaymentGateway, emailGateway gateway.EmailGateway) *PaymentService {
	return &PaymentService{
		paymentGateway,
		emailGateway,
	}
}

func (ps *PaymentService) CreatePix(amountInCents int) error {
	if amountInCents <= 0 {
		return ErrInvalidAmount
	}

	pixResponse, err := ps.gateway.CreatePixPayment(&PixPaymentRequest{
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
	_, err := ps.gateway.CheckPixPaymentStatus(&CheckPixPaymentStatusRequest{Id: id})

	return err
}
