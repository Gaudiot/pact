package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pact/core"
	student_repository "github.com/pact/core/database/repository/student"
	"github.com/pact/src/payment"
	"github.com/pact/src/user"
)

func main() {
	err := core.LoadEnvConfig()
	if err != nil {
		log.Fatalln(err)
	}

	router := chi.NewRouter()

	userRouter := user.NewUserRouter(
		student_repository.NewStub(),
	)
	router.Route("/payment", payment.PaymentRouter)
	router.Route("/user", userRouter.Router)

	log.Printf("Server is running on port :8000\n")
	http.ListenAndServe(":8000", router)
}
