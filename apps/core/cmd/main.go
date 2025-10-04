package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/capriolo/hackyeah2025/core/internal/static"
)

func main() {
	mux := http.NewServeMux()

	static.Serve(mux)

	mux.HandleFunc("GET /api/v1/test", func(w http.ResponseWriter, r *http.Request) {

		data := map[string]any{
			"version": 1,
			"name":    "HackYeah2025",
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	port := 80
	log.Printf("Serve api on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)

	if err != nil {
		panic(err)
	}
}
