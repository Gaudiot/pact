package payment

import (
	"errors"
	"testing"

	student_repository "github.com/pact/core/database/repository/student"
	"github.com/pact/src/gateway"
	payment_gateway "github.com/pact/src/gateway/payment"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	mockStudent = &student_repository.Student{
		ID:    "userId",
		Name:  "name",
		Email: "name@email.com",
	}
)

func setupTest(t *testing.T) (
	*PaymentService,
	*student_repository.MockRepository,
	*payment_gateway.MockPaymentGateway,
	*gateway.MockEmailGateway,
) {
	t.Helper()

	mockStudentRepository := new(student_repository.MockRepository)
	mockPaymentGateway := new(payment_gateway.MockPaymentGateway)
	mockEmailGateway := new(gateway.MockEmailGateway)

	sut := NewPaymentService(mockStudentRepository, mockPaymentGateway, mockEmailGateway)

	return sut, mockStudentRepository, mockPaymentGateway, mockEmailGateway
}

func TestCreatePix(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, mockStudentRepository, mockPaymentGateway, mockEmailGateway := setupTest(t)
		expected := &payment_gateway.PixPaymentResponse{
			Id:                  "pix-123",
			AmountInCents:       1000,
			PaymentCode:         "code",
			PaymentQRCodeBase64: "qrcode",
			CreatedAt:           "now",
			ExpiresAt:           "tomorrow",
		}
		mockStudentRepository.On("GetByID", mock.Anything, mock.Anything).Return(mockStudent, nil).Once()
		mockPaymentGateway.On("CreatePixPayment", mock.Anything).Return(expected, nil).Once()
		mockEmailGateway.On("SendPixPaymentEmail", mock.Anything).Return(nil).Once()

		err := sut.CreatePix("userId", 1000)

		mockStudentRepository.AssertExpectations(t)
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
	t.Run("User Repository Error", func(t *testing.T) {
		tests := []struct {
			testName      string
			expectedError error
		}{
			{
				"User Not Found",
				student_repository.ErrStudentNotFound,
			},
			{
				"Repository Internal Error",
				errors.New("internal_error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				sut, mockStudentRepository, mockPaymentGateway, mockEmailGateway := setupTest(t)
				mockStudentRepository.On("GetByID", mock.Anything, mock.Anything).Return(nil, tt.expectedError).Once()

				err := sut.CreatePix("userId", 1000)

				require.ErrorIs(t, err, tt.expectedError)

				mockStudentRepository.AssertExpectations(t)
				mockPaymentGateway.AssertNumberOfCalls(t, "CreatePixPayment", 0)
				mockEmailGateway.AssertNumberOfCalls(t, "SendPixPaymentEmail", 0)
			})
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

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				sut, mockStudentRepository, mockPaymentGateway, mockEmailGateway := setupTest(t)
				err := sut.CreatePix("userId", tt.amountInCents)

				require.ErrorIs(t, err, ErrInvalidAmount)

				mockStudentRepository.AssertNumberOfCalls(t, "GetByID", 0)
				mockPaymentGateway.AssertNumberOfCalls(t, "CreatePixPayment", 0)
				mockEmailGateway.AssertNumberOfCalls(t, "SendPixPaymentEmail", 0)
			})
		}
	})
	t.Run("Payment Gateway Error", func(t *testing.T) {
		sut, mockStudentRepository, mockPaymentGateway, mockEmailGateway := setupTest(t)
		expectedError := errors.New("payment_gateway_error")
		mockStudentRepository.On("GetByID", mock.Anything, mock.Anything).Return(mockStudent, nil).Once()
		mockPaymentGateway.On("CreatePixPayment", mock.Anything).Return(nil, expectedError).Once()

		err := sut.CreatePix("userId", 1000)

		mockStudentRepository.AssertExpectations(t)
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
		sut, mockStudentRepository, mockPaymentGateway, mockEmailGateway := setupTest(t)
		expected := &payment_gateway.PixPaymentResponse{
			Id:                  "pix-123",
			AmountInCents:       1000,
			PaymentCode:         "code",
			PaymentQRCodeBase64: "qrcode",
			CreatedAt:           "now",
			ExpiresAt:           "tomorrow",
		}
		expectedError := errors.New("email_gateway_error")
		mockStudentRepository.On("GetByID", mock.Anything, mock.Anything).Return(mockStudent, nil).Once()
		mockPaymentGateway.On("CreatePixPayment", mock.Anything).Return(expected, nil).Once()
		mockEmailGateway.On("SendPixPaymentEmail", mock.Anything).Return(expectedError).Once()

		err := sut.CreatePix("userId", 1000)

		mockStudentRepository.AssertExpectations(t)
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
		sut, _, mockPaymentGateway, _ := setupTest(t)

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
		sut, _, mockPaymentGateway, _ := setupTest(t)
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
