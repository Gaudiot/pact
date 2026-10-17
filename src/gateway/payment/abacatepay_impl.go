package payment_gateway

import (
	"context"
	"time"

	"github.com/AbacatePay/abacatepay-go-sdk/v1"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/pix"
	v1 "github.com/AbacatePay/go-types/v1"
	"github.com/pact/core"
)

const (
	CLIENT_TIMEOUT      = 10 * time.Second
	PIX_PAYMENT_TIMEOUT = 15 * time.Second
)

var _ PaymentGateway = &abacatepayPaymentGatewayImpl{}

type abacatepayPaymentGatewayImpl struct {
	client *abacatepay.Client
}

func NewAbacatepayPaymentGatewayImpl() (PaymentGateway, error) {
	apiKey := core.GetEnvConfig().AbacatePayAPIKey()
	client, err := abacatepay.New(abacatepay.ClientConfig{
		APIKey:  apiKey,
		Timeout: CLIENT_TIMEOUT,
	})
	if err != nil {
		return nil, err
	}

	return &abacatepayPaymentGatewayImpl{
		client: client,
	}, nil
}

func (pg *abacatepayPaymentGatewayImpl) CreatePixPayment(request *PixPaymentRequest) (*PixPaymentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PIX_PAYMENT_TIMEOUT)
	defer cancel()

	response, err := pg.client.PIX.Create(
		ctx,
		pix.CreateQrCodePIXBody{
			Amount:    request.AmountInCents,
			ExpiresIn: &request.ExpiresInSeconds,
		},
	)
	if err != nil {
		return nil, err
	}

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

func (pg *abacatepayPaymentGatewayImpl) CheckPixPaymentStatus(request *CheckPixPaymentStatusRequest) (*CheckPixPaymentStatusResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PIX_PAYMENT_TIMEOUT)
	defer cancel()

	result, err := pg.client.PIX.Status(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	var status PaymentStatus
	switch result.Data.Status {
	case v1.PaymentStatusPending:
		status = Pending
	case v1.PaymentStatusPaid:
		status = Paid
	case v1.PaymentStatusExpired:
		status = Expired
	default:
		status = Other
	}

	return &CheckPixPaymentStatusResponse{
		Status: status,
	}, nil
}
