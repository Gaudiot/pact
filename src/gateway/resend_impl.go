package gateway

import (
	"fmt"

	"github.com/pact/core"
	"github.com/resend/resend-go/v4"
)

var _ EmailGateway = &resendEmailGatewayImpl{}

type resendEmailGatewayImpl struct {
	client *resend.Client
}

func NewResendEmailGateway() EmailGateway {
	resendAPIKey := core.GetEnvConfig().ResendAPIKey()
	client := resend.NewClient(resendAPIKey)
	return &resendEmailGatewayImpl{
		client: client,
	}
}

func (eg *resendEmailGatewayImpl) SendPixPaymentEmail(args pixSendPixPaymentEmailArgs) error {
	amount := fmt.Sprintf("R$%d,%02d", args.AmountInCents/100, args.AmountInCents%100)
	expirationDate := args.ExpirationDate.Format("02/01/2006")

	params := &resend.SendEmailRequest{
		From:    "CTCL <ctcl-noreply@gaudiot.com>",
		To:      []string{args.ReceiverEmail},
		Subject: "Pix Payment",
		Html:    "<p>Pix Payment {" + amount + "} {companyName} {" + expirationDate + "} {pixCode} {pixQRCode}</p>",
	}

	sent, err := eg.client.Emails.Send(params)
	if err != nil {
		return err
	}

	fmt.Println(sent.Id)
	return nil
}
