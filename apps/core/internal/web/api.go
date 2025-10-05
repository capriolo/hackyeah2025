package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/capriolo/hackyeah2025/core/internal/module/day"
	"github.com/capriolo/hackyeah2025/core/internal/store"
)

func ServeApi(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/v1/day", func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")

		res := day.Day(date)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("POST /api/v1/push", func(w http.ResponseWriter, r *http.Request) {
		var req store.PushSubscription

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := store.Read()
		data.AddUserSubscription("pio", req)
		store.Write(data)

		log.Println("Zasubskrybowano")

		SendNotification(req, []byte(`{"title":"Witaj","body":"Działa Web Push!"}`))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	})

	mux.HandleFunc("GET /api/v1/push", func(w http.ResponseWriter, r *http.Request) {
		data := store.Read()
		subs := data.UserSubscriptions("pio")

		for _, sub := range subs {
			SendNotification(sub, []byte(`{"title":"Testowe powiadomienie","body":"Działa Web Push!"}`))
		}
	})

	mux.HandleFunc("GET /api/v1", func(w http.ResponseWriter, r *http.Request) {

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

}
