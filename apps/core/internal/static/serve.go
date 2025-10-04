package static

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed webapp/*
var static embed.FS

func Serve(mux *http.ServeMux) {
	webappFS, err := fs.Sub(static, "webapp")
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(webappFS))

	mux.Handle("GET /", fileServer)
}
