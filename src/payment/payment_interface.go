package payment

// MARK: - Models

type PixPaymentRequest struct {
	AmountInCents int
}

type PixPaymentResponse struct {
	Id                  string
	AmountInCents       int
	PaymentCode         string
	PaymentQRCodeBase64 string
	CreatedAt           string
	ExpiresAt           string
}

// MARK: - Interface

type PaymentService interface {
	CreatePixPayment(request *PixPaymentRequest) (*PixPaymentResponse, error)
}
