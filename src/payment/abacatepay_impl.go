package payment

import (
	"context"
	"time"

	"github.com/AbacatePay/abacatepay-go-sdk/v1"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/pix"
)

var _ PaymentService = &PaymentServiceImpl{}

type PaymentServiceImpl struct{}

const (
	CLIENT_TIMEOUT      = 10 * time.Second
	PIX_PAYMENT_TIMEOUT = 15 * time.Second
)

func (ps *PaymentServiceImpl) CreatePixPayment(request *PixPaymentRequest) (*PixPaymentResponse, error) {
	client, err := abacatepay.New(abacatepay.ClientConfig{
		APIKey:  "abc_dev_nzDCTjHsp0P0WP3Xj3ZXdhHM",
		Timeout: CLIENT_TIMEOUT,
	})
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), PIX_PAYMENT_TIMEOUT)
	defer cancel()

	response, err := client.PIX.Create(
		ctx,
		pix.CreateQrCodePIXBody{
			Amount: request.AmountInCents,
		},
	)

	data := response.Data
	pixResponse := PixPaymentResponse{
		Id:                  data.ID,
		AmountInCents:       data.Amount,
		PaymentCode:         data.BrCode,
		PaymentQRCodeBase64: data.BrCodeBase64,
		CreatedAt:           data.CreatedAt,
		ExpiresAt:           data.ExpiresAt,
	}

	return &pixResponse, nil
}
