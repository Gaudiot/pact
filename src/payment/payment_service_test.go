package payment

import (
	"errors"
	"testing"

	"github.com/pact/src/gateway"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) (
	*PaymentService,
	*MockPaymentGateway,
	*gateway.MockEmailGateway,
) {
	t.Helper()

	paymentGateway := new(MockPaymentGateway)
	emailGateway := new(gateway.MockEmailGateway)

	sut := NewPaymentService(paymentGateway, emailGateway)

	return sut, paymentGateway, emailGateway
}

func TestCreatePix(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, mockPaymentGateway, mockEmailGateway := setupTest(t)
		expected := &PixPaymentResponse{
			Id:                  "pix-123",
			AmountInCents:       1000,
			PaymentCode:         "code",
			PaymentQRCodeBase64: "qrcode",
			CreatedAt:           "now",
			ExpiresAt:           "tomorrow",
		}
		mockPaymentGateway.On("CreatePixPayment", mock.Anything).Return(expected, nil).Once()
		mockEmailGateway.On("SendPixPaymentEmail", mock.Anything).Return(nil).Once()

		err := sut.CreatePix(1000)

		mockPaymentGateway.AssertExpectations(t)
		mockEmailGateway.AssertExpectations(t)

		if !errors.Is(err, nil) {
			t.Errorf(
				"got error %v, want %v",
				err,
				nil,
			)
		}
	})
	t.Run("Invalid Amount", func(t *testing.T) {
		tests := []struct {
			testName      string
			amountInCents int
		}{
			{
				"zero value",
				0,
			},
			{
				"negative value",
				-1000,
			},
		}
		var _ = tests

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				sut, mockPaymentGateway, mockEmailGateway := setupTest(t)
				err := sut.CreatePix(tt.amountInCents)

				require.ErrorIs(t, err, ErrInvalidAmount)

				mockPaymentGateway.AssertNumberOfCalls(t, "CreatePixPayment", 0)
				mockEmailGateway.AssertNumberOfCalls(t, "SendPixPaymentEmail", 0)
			})
		}
	})
	t.Run("Payment Gateway Error", func(t *testing.T) {
		sut, mockPaymentGateway, mockEmailGateway := setupTest(t)
		expectedError := errors.New("payment_gateway_error")
		mockPaymentGateway.On("CreatePixPayment", mock.Anything).Return(nil, expectedError).Once()

		err := sut.CreatePix(1000)

		mockPaymentGateway.AssertExpectations(t)
		mockEmailGateway.AssertNumberOfCalls(t, "SendPixPaymentEmail", 0)

		if !errors.Is(err, expectedError) {
			t.Errorf(
				"got error %v, want %v",
				err,
				expectedError,
			)
		}
	})
	t.Run("Email Gateway Error", func(t *testing.T) {
		sut, mockPaymentGateway, mockEmailGateway := setupTest(t)
		expected := &PixPaymentResponse{
			Id:                  "pix-123",
			AmountInCents:       1000,
			PaymentCode:         "code",
			PaymentQRCodeBase64: "qrcode",
			CreatedAt:           "now",
			ExpiresAt:           "tomorrow",
		}
		expectedError := errors.New("email_gateway_error")
		mockPaymentGateway.On("CreatePixPayment", mock.Anything).Return(expected, nil).Once()
		mockEmailGateway.On("SendPixPaymentEmail", mock.Anything).Return(expectedError).Once()

		err := sut.CreatePix(1000)

		mockPaymentGateway.AssertExpectations(t)
		mockEmailGateway.AssertExpectations(t)

		if !errors.Is(err, expectedError) {
			t.Errorf(
				"got error %v, want %v",
				err,
				expectedError,
			)
		}
	})
}

func TestCheckPixStatus(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, mockPaymentGateway, _ := setupTest(t)

		mockPaymentGateway.On("CheckPixPaymentStatus", mock.Anything).Return(nil, nil)
		err := sut.CheckPixStatus("1")

		mockPaymentGateway.AssertExpectations(t)

		if !errors.Is(err, nil) {
			t.Errorf(
				"got error %v, want %v",
				err,
				nil,
			)
		}
	})
	t.Run("Payment Gateway Error", func(t *testing.T) {
		sut, mockPaymentGateway, _ := setupTest(t)
		expectedError := errors.New("payment_gateway_error")

		mockPaymentGateway.On("CheckPixPaymentStatus", mock.Anything).Return(nil, expectedError)
		err := sut.CheckPixStatus("1")

		mockPaymentGateway.AssertExpectations(t)

		if !errors.Is(err, expectedError) {
			t.Errorf(
				"got error %v, want %v",
				err,
				expectedError,
			)
		}
	})
}
