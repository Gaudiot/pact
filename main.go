package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pact/core"
	"github.com/pact/src/payment"
)

func main() {
	err := core.LoadEnvConfig()
	if err != nil {
		log.Fatalln(err)
	}

	router := chi.NewRouter()

	router.Route("/payment", payment.PaymentRouter)

	log.Printf("Server is running on port :8000\n")
	http.ListenAndServe(":8000", router)
}
