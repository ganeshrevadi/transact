package main

import (
	"encoding/json"
	"log"
	"net/http"

	"payment-system/internal/payment"
)

func main(){
	mux := http.NewServeMux()
	count := 0

	paymentHandler := payment.NewHandler(payment.NewService())
	mux.HandleFunc("POST /payments", paymentHandler.Create)
	mux.HandleFunc("GET /payments/{id}", paymentHandler.Get)

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", 	"application/json")
		w.WriteHeader(http.StatusOK)


		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
		count++
		log.Println("Pinged health endpoint: ",count )
	})

	log.Println("Payment API listening on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)

	}
}
