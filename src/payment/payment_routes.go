package payment

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func PaymentRouter(r chi.Router) {
	r.Get("/status/{ID}", handleGetPaymentStatus)
	r.Post("/", handleRoot)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var paymentGateway PaymentGateway = &AbacatepayPaymentGatewayImpl{}
	response, err := paymentGateway.CreatePixPayment(&PixPaymentRequest{AmountInCents: req.AmountInCents})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(strconv.Itoa(response.AmountInCents) + "|%%|" + response.PaymentCode + " |%%| " + response.PaymentQRCodeBase64))
}

func handleGetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	paymentId := chi.URLParam(r, "ID")
	var _ = paymentId
}

// MARK: - Improve later

type PaymentRequest struct {
	AmountInCents int `json:"amount_in_cents"`
}
