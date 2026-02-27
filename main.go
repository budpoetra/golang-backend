package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next(w, r)

		log.Printf(
			"Method: %s | Path: %s | Duration: %v",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]string{
			"app_name": "My Golang Backend",
			"version":  "1.0.0",
		},
	}

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", loggerMiddleware(handleIndex))

	port := ":8080"
	log.Printf("Server berjalan di http://localhost%s", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
