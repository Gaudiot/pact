package payment_gateway

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

type PaymentStatus string

const (
	Pending PaymentStatus = "PENDING"
	Paid    PaymentStatus = "PAID"
	Expired PaymentStatus = "EXPIRED"
	Other   PaymentStatus = "OTHER"
)

type CheckPixPaymentStatusRequest struct {
	Id string
}

type CheckPixPaymentStatusResponse struct {
	Status PaymentStatus
}

// MARK: - Interface

type PaymentGateway interface {
	CreatePixPayment(request *PixPaymentRequest) (*PixPaymentResponse, error)
	CheckPixPaymentStatus(request *CheckPixPaymentStatusRequest) (*CheckPixPaymentStatusResponse, error)
}
