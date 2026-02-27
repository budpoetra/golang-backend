package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Response adalah struktur standar untuk API kita
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// loggerMiddleware berfungsi untuk mencatat setiap request yang masuk
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Lanjutkan ke handler utama
		next(w, r)
		
		// Cetak log setelah request selesai diproses
		log.Printf(
			"Method: %s | Path: %s | Duration: %v",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Set header agar response berupa JSON
	w.Header().Set("Content-Type", "application/json")

	// Data yang ingin dikirim
	res := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]string{
			"app_name": "My Golang Backend",
			"version":  "1.0.0",
		},
	}

	// Mengubah struct menjadi JSON dan mengirimnya
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w.Encode(res))
}

func main() {
	// Routing
	http.HandleFunc("/", loggerMiddleware(handleIndex))

	// Menjalankan server
	port := ":8080"
	log.Printf("Server berjalan di http://localhost%s", port)
	
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
