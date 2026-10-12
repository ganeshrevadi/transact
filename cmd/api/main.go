package main

import (
	"net/http"
	"log"
	"encoding/json"
)

func main(){
	mux := http.NewServeMux()
	count := 0

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
