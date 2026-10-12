package payment

import "errors"

type PaymentService struct {
	gateway PaymentGateway
}

func NewPaymentService(paymentGateway PaymentGateway) *PaymentService {
	return &PaymentService{
		gateway: paymentGateway,
	}
}

func (ps *PaymentService) CreatePix(amountInCents int) error {
	if amountInCents <= 0 {
		return errors.New("Amount is not a positive number.")
	}

	_, err := ps.gateway.CreatePixPayment(&PixPaymentRequest{
		AmountInCents: amountInCents,
	})
	if err != nil {
		return err
	}

	return nil
}
