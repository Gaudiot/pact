package payment_gateway

import "github.com/stretchr/testify/mock"

type MockPaymentGateway struct {
	mock.Mock
}

func (m *MockPaymentGateway) CreatePixPayment(request *PixPaymentRequest) (*PixPaymentResponse, error) {
	args := m.Called(request)

	response, _ := args.Get(0).(*PixPaymentResponse)

	return response, args.Error(1)
}

func (m *MockPaymentGateway) CheckPixPaymentStatus(request *CheckPixPaymentStatusRequest) (*CheckPixPaymentStatusResponse, error) {
	args := m.Called(request)

	response, _ := args.Get(0).(*CheckPixPaymentStatusResponse)

	return response, args.Error(1)
}
