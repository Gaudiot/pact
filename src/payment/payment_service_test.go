package payment

import (
	"errors"
	"testing"

	"github.com/pact/src/gateway"
	"github.com/stretchr/testify/mock"
)

var (
	mockPaymentGateway = &MockPaymentGateway{}
	mockEmailGateway   = &gateway.MockEmailGateway{}
)

func TestCreatePix(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

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

		sut := NewPaymentService(mockPaymentGateway, mockEmailGateway)

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
	t.Run("Invalid Amount", func(t *testing.T) {})
	t.Run("Payment Gateway Error", func(t *testing.T) {})
	t.Run("Email Gateway Error", func(t *testing.T) {})
}

func TestCheckPixStatus(t *testing.T) {
	t.Run("Success", func(t *testing.T) {})
	t.Run("Payment Gateway Error", func(t *testing.T) {})
}

// func TestCreatePix(t *testing.T) {
// 	tests := []struct {
// 		testName      string
// 		amountInCents int
// 		expectedErr   error
// 	}{
// 		{
// 			testName:      "success",
// 			amountInCents: 1000,
// 			expectedErr:   nil,
// 		},
// 		{
// 			testName:      "zero amount",
// 			amountInCents: 0,
// 			expectedErr:   nil,
// 		},
// 		{
// 			testName:      "negative amount",
// 			amountInCents: -1000,
// 			expectedErr:   nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.testName, func(t *testing.T) {
// 			paymentGateway := &MockPaymentGateway{}
// 			mockEmailGateway := &gateway.MockEmailGateway{}
// 			service := NewPaymentService(paymentGateway, mockEmailGateway)

// 			request := &PixPaymentRequest{
// 				AmountInCents: 1000,
// 			}

// 			expected := &PixPaymentResponse{
// 				Id:            "pix-123",
// 				AmountInCents: 1000,
// 			}

// 			paymentGateway.On("CreatePixPayment", request).
// 				Return(expected, nil).
// 				Once()

// 			mockEmailGateway.On("SendPixPaymentEmail").
// 				Return(nil).
// 				Once()

// 			err := service.CreatePix(tt.amountInCents)

// 			paymentGateway.AssertExpectations(t)

// 			if !errors.Is(err, tt.expectedErr) {
// 				t.Errorf(
// 					"got error %v, want %v",
// 					err,
// 					tt.expectedErr,
// 				)
// 			}
// 		})
// 	}
// }
