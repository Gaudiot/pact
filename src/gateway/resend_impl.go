package gateway

import (
	"encoding/base64"
	"fmt"
	"strings"

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

func (eg *resendEmailGatewayImpl) SendPixPaymentEmail(args *PixSendPixPaymentEmailArgs) error {
	amount := fmt.Sprintf("R$%d,%02d", args.AmountInCents/100, args.AmountInCents%100)
	expirationDate := args.ExpirationDate.Format("02/01/2006")

	qrCodeBase64 := args.PixQRCode
	if index := strings.Index(qrCodeBase64, ","); index != -1 {
		qrCodeBase64 = qrCodeBase64[index+1:]
	}

	qrCodeBytes, err := base64.StdEncoding.DecodeString(
		strings.TrimSpace(qrCodeBase64),
	)
	if err != nil {
		return err
	}

	params := &resend.SendEmailRequest{
		From:    "CTCL <ctcl-noreply@gaudiot.com>",
		To:      []string{args.ReceiverEmail},
		Subject: "Pix Payment",
		Template: &resend.EmailTemplate{
			Id: "pact_pix_payment",
			Variables: map[string]any{
				"name":            args.UserName,
				"amount":          amount,
				"company_name":    args.CompanyName,
				"expiration_date": expirationDate,
				"pix_code":        args.PixCode,
			},
		},
		Attachments: []*resend.Attachment{
			{
				Filename:    "pix.png",
				Content:     qrCodeBytes,
				ContentType: "image/png",
				ContentId:   "pix-qrcode",
			},
		},
	}

	sent, err := eg.client.Emails.Send(params)
	if err != nil {
		return err
	}

	fmt.Println(sent.Id)
	return nil
}
