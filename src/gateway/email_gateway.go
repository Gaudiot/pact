package gateway

import "time"

type PixSendPixPaymentEmailArgs struct {
	UserName       string
	AmountInCents  int
	CompanyName    string
	ExpirationDate time.Time
	PixCode        string
	PixQRCode      string
	ReceiverEmail  string
}

type EmailGateway interface {
	SendPixPaymentEmail(args *PixSendPixPaymentEmailArgs) error
}
