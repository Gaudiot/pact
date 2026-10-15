package gateway

import "github.com/stretchr/testify/mock"

type MockEmailGateway struct {
	mock.Mock
}

func (m *MockEmailGateway) SendPixPaymentEmail(request *PixSendPixPaymentEmailArgs) error {
	args := m.Called(request)

	return args.Error(0)
}
