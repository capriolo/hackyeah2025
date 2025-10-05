package web

import (
	"log"
	"os"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/capriolo/hackyeah2025/core/internal/store"
)

func SendNotification(sub store.PushSubscription, payload []byte) {
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
