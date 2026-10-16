package payment

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	student_repository "github.com/pact/core/database/repository/student"
	"github.com/pact/src/gateway"
	payment_gateway "github.com/pact/src/gateway/payment"
)

func PaymentRouter(r chi.Router) {
	r.Get("/status/{ID}", handleGetPaymentStatus)
	r.Post("/", handleRoot)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	studentRepository, err := student_repository.NewPgxStudentRepository("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	paymentGateway, err := payment_gateway.NewAbacatepayPaymentGatewayImpl()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	emailGateway := gateway.NewResendEmailGateway()
	paymentService := NewPaymentService(studentRepository, paymentGateway, emailGateway)

	err = paymentService.CreatePix(req.StudentId, req.AmountInCents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleGetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	paymentId := chi.URLParam(r, "ID")
	var _ = paymentId
}

// MARK: - Improve later

type PaymentRequest struct {
	StudentId     string `json:"student_id"`
	AmountInCents int    `json:"amount_in_cents"`
}
