package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/capriolo/hackyeah2025/core/internal/platform/envloader"
	"github.com/capriolo/hackyeah2025/core/internal/static"
	"github.com/capriolo/hackyeah2025/core/internal/web"
)

func main() {
	mux := http.NewServeMux()

	web.ServeApi(mux)
	static.Serve(mux)

	var handler http.Handler = mux
	handler = web.CorsMiddleware(handler)

	port := 80
	log.Printf("Serve api on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)

	if err != nil {
		panic(err)
	}
}
