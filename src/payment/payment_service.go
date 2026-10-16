package payment

import (
	"context"
	"errors"
	"time"

	student_repository "github.com/pact/core/database/repository/student"
	"github.com/pact/src/gateway"
	payment_gateway "github.com/pact/src/gateway/payment"
)

var (
	ErrInvalidAmount  = errors.New("amount must be greater than zero")
	pixExpirationTime = time.Hour * 24 * 7
)

type PaymentService struct {
	studentRepository student_repository.Repository
	paymentGateway    payment_gateway.PaymentGateway
	emailGateway      gateway.EmailGateway
}

func NewPaymentService(studentRepository student_repository.Repository, paymentGateway payment_gateway.PaymentGateway, emailGateway gateway.EmailGateway) *PaymentService {
	return &PaymentService{
		studentRepository,
		paymentGateway,
		emailGateway,
	}
}

func (ps *PaymentService) CreatePix(userId string, amountInCents int) error {
	if amountInCents <= 0 {
		return ErrInvalidAmount
	}
	ctx := context.Background()

	student, err := ps.studentRepository.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	pixResponse, err := ps.paymentGateway.CreatePixPayment(&payment_gateway.PixPaymentRequest{
		AmountInCents: amountInCents,
	})
	if err != nil {
		return err
	}

	err = ps.emailGateway.SendPixPaymentEmail(&gateway.PixSendPixPaymentEmailArgs{
		UserName:       student.Name,
		AmountInCents:  pixResponse.AmountInCents,
		CompanyName:    "CTCL",
		ExpirationDate: time.Now().Add(pixExpirationTime),
		PixCode:        pixResponse.PaymentCode,
		PixQRCode:      pixResponse.PaymentQRCodeBase64,
		ReceiverEmail:  student.Email,
	})
	if err != nil {
		return err
	}

	return nil
}

func (ps *PaymentService) CheckPixStatus(id string) error {
	_, err := ps.paymentGateway.CheckPixPaymentStatus(&payment_gateway.CheckPixPaymentStatusRequest{Id: id})

	return err
}
