package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/SherClockHolmes/webpush-go"
	_ "github.com/capriolo/hackyeah2025/core/internal/platform/envloader"
	"github.com/capriolo/hackyeah2025/core/internal/static"
	"github.com/capriolo/hackyeah2025/core/internal/web"
)

type subscriptionReqKeys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}
type subscriptionReq struct {
	Endpoint string `json:"endpoint"`
	// ExpirationTime string              `json:"expirationTime,omitempty"`
	Keys subscriptionReqKeys `json:"keys"`
}

func main() {
	mux := http.NewServeMux()

	static.Serve(mux)

	// generateCompatibleKeys()
	debugKeys()

	var mu sync.RWMutex
	var subs []subscriptionReq

	mux.HandleFunc("POST /api/v1/push", func(w http.ResponseWriter, r *http.Request) {
		var req subscriptionReq

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mu.Lock()
		defer mu.Unlock()
		subs = append(subs, req)

		log.Println("Zasubskrybowano")

		SendNotification(req, []byte(`{"title":"Witaj","body":"Działa Web Push!"}`))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	})

	mux.HandleFunc("GET /api/v1/push", func(w http.ResponseWriter, r *http.Request) {

		mu.RLock()

		for _, sub := range subs {
			SendNotification(sub, []byte(`{"title":"Testowe powiadomienie","body":"Działa Web Push!"}`))
		}
		defer mu.RUnlock()
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

	var handler http.Handler = mux
	handler = web.CorsMiddleware(handler)

	port := 80
	log.Printf("Serve api on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)

	if err != nil {
		panic(err)
	}
}

func SendNotification(sub subscriptionReq, payload []byte) {
	resp, err := webpush.SendNotification(payload, &webpush.Subscription{
		Endpoint: sub.Endpoint,
		Keys: webpush.Keys{
			Auth:   sub.Keys.Auth,
			P256dh: sub.Keys.P256dh,
		},
	}, &webpush.Options{
		VAPIDPublicKey:  os.Getenv("VAPID_PUBLIC_KEY"),
		VAPIDPrivateKey: os.Getenv("VAPID_PRIVATE_KEY"),
		TTL:             60,
	})

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println("Powiadomienie wysłane:", resp.Status)
}

func debugKeys() {
	publicKey := os.Getenv("VAPID_PUBLIC_KEY")
	privateKey := os.Getenv("VAPID_PRIVATE_KEY")

	fmt.Printf("Public Key: %s\n", publicKey)
	fmt.Printf("Private Key: %s\n", privateKey)

	// Sprawdź długość kluczy po zdekodowaniu
	publicKeyBytes, err := base64.RawURLEncoding.DecodeString(publicKey)
	if err != nil {
		log.Fatal("Error decoding public key:", err)
	}

	privateKeyBytes, err := base64.RawURLEncoding.DecodeString(privateKey)
	if err != nil {
		log.Fatal("Error decoding private key:", err)
	}

	fmt.Printf("Public key length: %d bytes\n", len(publicKeyBytes))
	fmt.Printf("Private key length: %d bytes\n", len(privateKeyBytes))
	fmt.Printf("Public key first byte: 0x%02x\n", publicKeyBytes[0])

	// Dla P-256 klucz publiczny powinien mieć 65 bajtów (uncompressed)
	if len(publicKeyBytes) != 65 {
		log.Fatalf("ERROR: Public key should be 65 bytes, but got %d", len(publicKeyBytes))
	}

	if publicKeyBytes[0] != 0x04 {
		log.Fatalf("ERROR: Public key should start with 0x04, but starts with 0x%02x", publicKeyBytes[0])
	}

	fmt.Println("✓ Keys appear to be in correct format")
}

func generateCompatibleKeys() {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("VAPID_PUBLIC_KEY=%s\n", publicKey)
	fmt.Printf("VAPID_PRIVATE_KEY=%s\n", privateKey)
}
