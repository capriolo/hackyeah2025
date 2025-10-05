package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/capriolo/hackyeah2025/core/internal/module/day"
	"github.com/capriolo/hackyeah2025/core/internal/module/sync"
	"github.com/capriolo/hackyeah2025/core/internal/store"
)

type userProfile struct {
	Name      string `json:"name"`
	CalGoogle bool   `json:"calGoogle"`
	Garmin    bool   `json:"garmin"`
}

func ServeApi(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/v1/day", func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")

		data := store.Read()

		user := "piotr"
		userData := &store.User{}

		for _, u := range data.Users {
			if u.Name == user {
				userData = &u
			}
		}

		res := day.Day(userData, date)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("GET /api/v1/profile", func(w http.ResponseWriter, r *http.Request) {
		data := store.Read()
		user := data.UserProfile("piotr")

		userProfile := userProfile{
			Name:      user.Name,
			CalGoogle: user.CalendarUrl != "",
			Garmin:    user.RunalyzeToken != "",
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(userProfile); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("POST /api/v1/sync/garmin", func(w http.ResponseWriter, r *http.Request) {
		sync.SyncGarmin("piotr")
	})

	mux.HandleFunc("POST /api/v1/sync/ical", func(w http.ResponseWriter, r *http.Request) {
		sync.SyncIcal("piotr")
	})

	mux.HandleFunc("POST /api/v1/push", func(w http.ResponseWriter, r *http.Request) {
		var req store.PushSubscription

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := store.Read()
		data.AddUserSubscription("piotr", req)
		store.Write(data)

		log.Println("Zasubskrybowano")

		SendNotification(req, []byte(`{"title":"Witaj","body":"Działa Web Push!"}`))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	})

	mux.HandleFunc("GET /api/v1/push", func(w http.ResponseWriter, r *http.Request) {
		data := store.Read()
		subs := data.UserSubscriptions("piotr")

		for _, sub := range subs {
			SendNotification(sub, []byte(`{"title":"Testowe powiadomienie","body":"Działa Web Push!"}`))
		}
	})

	mux.HandleFunc("POST /api/v1/profile", func(w http.ResponseWriter, r *http.Request) {
		var req userProfile

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := store.Read()
		data.SaveUserProfile(store.User{
			Name: req.Name,
		})
		store.Write(data)
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
