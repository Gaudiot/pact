package payment

import (
	"errors"
	"testing"
)

func TestCreatePix(t *testing.T) {
	tests := []struct {
		testName      string
		amountInCents int
		expectedErr   error
	}{
		{
			testName:      "success",
			amountInCents: 1000,
			expectedErr:   nil,
		},
		{
			testName:      "zero amount",
			amountInCents: 0,
			expectedErr:   nil,
		},
		{
			testName:      "negative amount",
			amountInCents: -1000,
			expectedErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			gateway := &MockPaymentGateway{}
			service := NewPaymentService(gateway)

			request := &PixPaymentRequest{
				AmountInCents: 1000,
			}

			expected := &PixPaymentResponse{
				Id:            "pix-123",
				AmountInCents: 1000,
			}

			gateway.On("CreatePixPayment", request).
				Return(expected, nil).
				Once()

			err := service.CreatePix(tt.amountInCents)

			gateway.AssertExpectations(t)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf(
					"got error %v, want %v",
					err,
					tt.expectedErr,
				)
			}
		})
	}
}
