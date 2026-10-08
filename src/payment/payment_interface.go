package payment

// MARK: - Models

type PixPaymentRequest struct {
	AuthKey          string
	AmountInCents    int
	ExpiresInSeconds int
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
	CreatePixPayment() int
}
